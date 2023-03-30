package data

import (
	_imageData "alta-cookit-be/features/images/data"
	_imageModel "alta-cookit-be/features/images/models"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
	"alta-cookit-be/features/transaction_details"
	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"
	_userModel "alta-cookit-be/features/users/data"
)

func ConvertToGorm(entity *transaction_details.TransactionDetailEntity) *_transactionDetailModel.TransactionDetail {
	gorm := _transactionDetailModel.TransactionDetail{
		IngredientID: entity.IngredientID,
		Quantity:     entity.Quantity,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]transaction_details.TransactionDetailEntity) *[]_transactionDetailModel.TransactionDetail {
	gorms := []_transactionDetailModel.TransactionDetail{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_transactionDetailModel.TransactionDetail, userGorm *_userModel.User, recipeGorm *_recipeModel.Recipe, imageGorms *[]_imageModel.Image, ingredientGorm *_ingredientModel.Ingredient) *transaction_details.TransactionDetailEntity {
	entity := transaction_details.TransactionDetailEntity{
		ID:             gorm.ID,
		LoggedInUserID: gorm.ID,
		Quantity:       gorm.Quantity,
	}
	if userGorm != nil {
		entity.SellerUserID = userGorm.ID
		entity.SellerUsername = userGorm.Username
	}
	if recipeGorm != nil {
		entity.RecipeName = recipeGorm.Name
	}
	if imageGorms != nil {
		entity.RecipeImageEntities = *_imageData.ConvertToEntities(imageGorms)
	}
	if ingredientGorm != nil {
		entity.IngredientName = ingredientGorm.Name
		entity.Price = ingredientGorm.Price * float64(gorm.Quantity)
	}
	return &entity
}
