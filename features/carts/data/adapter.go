package data

import (
	"alta-cookit-be/features/carts"
	_cartModel "alta-cookit-be/features/carts/models"
	_imageModel "alta-cookit-be/features/images/models"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
	_userModel "alta-cookit-be/features/users/data"
	_imageData "alta-cookit-be/features/images/data"
)

func ConvertToGorm(entity *carts.CartEntity) *_cartModel.Cart {
	gorm := _cartModel.Cart{
		UserID:       entity.UserID,
		IngredientID: entity.IngredientID,
		Quantity:     entity.Quantity,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]carts.CartEntity) *[]_cartModel.Cart {
	gorms := []_cartModel.Cart{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_cartModel.Cart, userGorm *_userModel.User, recipeGorm *_recipeModel.Recipe, imageGorms *[]_imageModel.Image, ingredientGorm *_ingredientModel.Ingredient) *carts.CartEntity {
	entity := carts.CartEntity{
		ID:       gorm.ID,
		UserID:   gorm.ID,
		Quantity: gorm.Quantity,
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
