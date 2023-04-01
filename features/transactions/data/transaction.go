package data

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/transaction_details"
	"alta-cookit-be/features/transactions"
	"alta-cookit-be/features/users"
	"alta-cookit-be/utils/consts"
	"errors"
	"fmt"
	"strings"

	_transactionDetailData "alta-cookit-be/features/transaction_details/data"
	_transactionModel "alta-cookit-be/features/transactions/models"
	_userModel "alta-cookit-be/features/users/data"

	"gorm.io/gorm"
)

type TransactionData struct {
	db             *gorm.DB
	userData       users.UserData
	recipeData     recipes.RecipeData_
	imageData      images.ImageData_
	ingredientData ingredients.IngredientData_
}

func New(db *gorm.DB, userData users.UserData, recipeData recipes.RecipeData_, imageData images.ImageData_, ingredientData ingredients.IngredientData_) transactions.TransactionData_ {
	return &TransactionData{
		db:             db,
		userData:       userData,
		recipeData:     recipeData,
		imageData:      imageData,
		ingredientData: ingredientData,
	}
}

func (d *TransactionData) ActionValidator(id, customerUserId uint) bool {
	tempGorm := _transactionModel.Transaction{}
	d.db.Model(&tempGorm).Where("id = ? AND user_id = ?", id, customerUserId).Find(&tempGorm)

	return tempGorm.ID != 0
}

func (d *TransactionData) SelectTransactionById(id uint) *_transactionModel.Transaction {
	tempGorm := _transactionModel.Transaction{}

	d.db.Where("id = ?", id).Find(&tempGorm)

	return &tempGorm
}

func (d *TransactionData) SelectTransactionByTransactionDetailId(transactionDetailId uint) *_transactionModel.Transaction {
	tempGorm := _transactionModel.Transaction{}

	subQuery := d.db.Table("transaction_details").Where("id = ?", transactionDetailId).Select("transaction_id")
	d.db.Where("id IN (?)", subQuery).Find(&tempGorm)

	return &tempGorm
}

func (d *TransactionData) SelectTransactionsByUserId(entity *transactions.TransactionEntity) (*[]transactions.TransactionEntity, error) {
	gorms := []_transactionModel.Transaction{}

	qString := ""
	for key, val := range entity.ExtractedQueryParams {
		if qString != "" {
			qString += " AND "
		}
		if key == "name" {
			qString += fmt.Sprintf("%s LIKE %s%s%s ", key, "'%", val, "%'")
		} else {
			qString += fmt.Sprintf("%s = '%s'", key, val)
		}
	}

	tx := d.db.Preload("TransactionDetails").Where("user_id = ?", entity.CustomerUserId).Where(qString).Limit(entity.DataLimit).Offset(entity.DataOffset).Order("created_at desc").Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []transactions.TransactionEntity{}
	for _, gorm := range gorms {
		subEntities := []transaction_details.TransactionDetailEntity{}
		for _, subGorm := range gorm.TransactionDetails {
			recipeGorm := d.recipeData.SelectRecipeByIngredientId(subGorm.IngredientID)
			userEntity := d.userData.SelectUserById(users.Core{ID: recipeGorm.UserID})
			userGorm := _userModel.CoreToModel(*userEntity)
			imageGorms := d.imageData.SelectImagesByRecipeId(recipeGorm.ID)
			ingredientGorm := d.ingredientData.SelectIngredientById(subGorm.IngredientID)
			subEntities = append(subEntities, *_transactionDetailData.ConvertToEntity(&subGorm, &userGorm, recipeGorm, imageGorms, ingredientGorm))
		}
		entities = append(entities, *ConvertToEntity(&gorm, &subEntities))
	}

	return &entities, nil
}

func (d *TransactionData) InsertTransaction(entity *transactions.TransactionEntity) (*transactions.TransactionEntity, error) {
	gorm := ConvertToGorm(entity)
	if entity.ID != 0 {
		fmt.Println(entity.ID)
		d.db.Model(&gorm).Where("id = ?", entity.ID).Update("virtual_account_number", entity.VirtualAccountNumber)
		return nil, nil
	}
	subEntities := []transaction_details.TransactionDetailEntity{}
	for _, subGorm := range gorm.TransactionDetails {
		recipeGorm := d.recipeData.SelectRecipeByIngredientId(subGorm.IngredientID)
		userEntity := d.userData.SelectUserById(users.Core{ID: recipeGorm.UserID})
		userGorm := _userModel.CoreToModel(*userEntity)
		imageGorms := d.imageData.SelectImagesByRecipeId(recipeGorm.ID)
		ingredientGorm := d.ingredientData.SelectIngredientById(subGorm.IngredientID)
		subEntities = append(subEntities, *_transactionDetailData.ConvertToEntity(&subGorm, &userGorm, recipeGorm, imageGorms, ingredientGorm))
	}

	for _, subEntity := range subEntities {
		gorm.TotalPrice += subEntity.Price
	}
	gorm.TotalPrice += gorm.ShippingFee

	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "ingredient_id") {
			return nil, errors.New(consts.INGREDIENT_InvalidIngredient)
		}
		return nil, tx.Error
	}

	for index, _ := range subEntities {
		subEntities[index].ID = gorm.TransactionDetails[index].ID
	}

	return ConvertToEntity(gorm, &subEntities), nil
}

func (d *TransactionData) UpdateTransactionStatusById(entity *transactions.TransactionEntity) error {
	tx := d.db.Table("transactions").Where("id = ?", entity.ID).Update("status", entity.Status)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *TransactionData) UpdateTransactionStatusByMidtrans(entity *transactions.TransactionEntity) error {
	tx := d.db.Debug().Where("order_id = ?", entity.OrderID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
