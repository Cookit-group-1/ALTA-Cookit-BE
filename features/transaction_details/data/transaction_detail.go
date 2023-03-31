package data

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/transaction_details"
	"alta-cookit-be/features/transactions"
	"alta-cookit-be/features/users"

	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"
	_userModel "alta-cookit-be/features/users/data"

	"gorm.io/gorm"
)

type TransactionDetailData struct {
	db              *gorm.DB
	userData        users.UserData
	recipeData      recipes.RecipeData_
	imageData       images.ImageData_
	ingredientData  ingredients.IngredientData_
	transactionData transactions.TransactionData_
}

func New(db *gorm.DB, userData users.UserData, recipeData recipes.RecipeData_, imageData images.ImageData_, ingredientData ingredients.IngredientData_, transactionData transactions.TransactionData_) transaction_details.TransactionDetailData_ {
	return &TransactionDetailData{
		db:             db,
		userData:       userData,
		recipeData:     recipeData,
		imageData:      imageData,
		ingredientData: ingredientData,
		transactionData: transactionData,
	}
}

func (d *TransactionDetailData) ActionValidator(id, loggedInUserId uint) bool {
	tempGorm := _transactionDetailModel.TransactionDetail{}
	d.db.Model(&tempGorm).Where("id = ? AND user_id = ?", id, loggedInUserId).Find(&tempGorm)

	return tempGorm.ID != 0
}

func (d *TransactionDetailData) SelectTransactionDetailById(entity *transaction_details.TransactionDetailEntity) (*transaction_details.TransactionDetailEntity, error) {
	gorm := _transactionDetailModel.TransactionDetail{}

	tx := d.db.Where("id = ?", entity.ID).Find(&gorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	recipeGorm := d.recipeData.SelectRecipeByIngredientId(gorm.IngredientID)
	userEntity := d.userData.SelectUserById(users.Core{ID: recipeGorm.UserID})
	userGorm := _userModel.CoreToModel(*userEntity)
	imageGorms := d.imageData.SelectImagesByRecipeId(recipeGorm.ID)
	ingredientGorm := d.ingredientData.SelectIngredientById(gorm.IngredientID)
	return ConvertToEntity(&gorm, &userGorm, recipeGorm, imageGorms, ingredientGorm), nil
}
