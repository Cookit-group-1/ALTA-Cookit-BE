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

func (d *TransactionData) SelectTransactionByUserId(entity *transactions.TransactionEntity) (*[]transactions.TransactionEntity, error) {
	gorms := []_transactionModel.Transaction{}

	tx := d.db.Preload("transaction_details").Where("user_id = ?", entity.CustomerUserId).Limit(entity.DataLimit).Offset(entity.DataOffset).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []transactions.TransactionEntity{}
	for _, gorm := range gorms {
		subEntities := []transaction_details.TransactionDetailEntity{}
		for _, subGorm := range gorm.TransactionDetail {
			userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
			userGorm := _userModel.CoreToModel(*userEntity)
			recipeGorm := d.recipeData.SelectRecipeByIngredientId(gorm.IngredientID)
			imageGorms := d.imageData.SelectImagesByRecipeId(recipeGorm.ID)
			ingredientGorm := d.ingredientData.SelectIngredientById(gorm.IngredientID)
			subEntities = append(subEntities, *_transactionDetailData.ConvertToEntity(&subGorm, &userGorm, recipeGorm, imageGorms, ingredientGorm))
		}
		entities = append(entities, *ConvertToEntity(&gorm, &subEntities))
	}

	return &entities, nil
}

func (d *TransactionData) InsertTransaction(entity *transactions.TransactionEntity) (*transactions.TransactionEntity, error) {
	gorm := ConvertToGorm(entity)
	subEntities := []transaction_details.TransactionDetailEntity{}
	for _, subGorm := range gorm.TransactionDetail {
		userEntity := d.userData.SelectUserById(users.Core{ID: gorm.UserID})
		userGorm := _userModel.CoreToModel(*userEntity)
		recipeGorm := d.recipeData.SelectRecipeByIngredientId(gorm.IngredientID)
		imageGorms := d.imageData.SelectImagesByRecipeId(recipeGorm.ID)
		ingredientGorm := d.ingredientData.SelectIngredientById(gorm.IngredientID)
		subEntities = append(subEntities, *_transactionDetailData.ConvertToEntity(&subGorm, &userGorm, recipeGorm, imageGorms, ingredientGorm))
	}

	for _, subEntity := range subEntities {
		gorm.TotalPrice += subEntity.Price
	}

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

	return ConvertToEntity(gorm, &subEntities), nil
}

func (d *TransactionData) UpdateTransactionById(entity *transactions.TransactionEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
