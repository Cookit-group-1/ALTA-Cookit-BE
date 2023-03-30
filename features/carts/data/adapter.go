package data

import (
	"alta-cookit-be/features/carts"
	_cartModel "alta-cookit-be/features/carts/models"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
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

func ConvertToEntity(gorm *_cartModel.Cart, recipeGorm *_recipeModel.Recipe, ingredientGorm *_ingredientModel.Ingredient) *carts.CartEntity {
	entity := carts.CartEntity{
		ID:       gorm.ID,
		UserID:   gorm.ID,
		Quantity: gorm.Quantity,
	}
	if recipeGorm != nil {
		entity.RecipeName = recipeGorm.Name
	}
	if ingredientGorm != nil {
		entity.IngredientName = ingredientGorm.Name
		entity.Price = ingredientGorm.Price * float64(gorm.Quantity)
	}
	return &entity
}
