package data

import (
	"alta-cookit-be/features/ingredients"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
)

func ConvertToGorm(entity *ingredients.IngredientEntity) *_ingredientModel.Ingredient {
	gorm := _ingredientModel.Ingredient{
		RecipeID: entity.RecipeID,
		Name:     entity.Name,
		Price:    entity.Price,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToEntity(gorm *_ingredientModel.Ingredient) *ingredients.IngredientEntity {
	return &ingredients.IngredientEntity{
		ID: gorm.ID,
		RecipeID: gorm.RecipeID,
		Name: gorm.Name,
		Price: gorm.Price,
	}
}
