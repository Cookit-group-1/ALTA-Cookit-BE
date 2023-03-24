package data

import (
	"alta-cookit-be/features/ingredients"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_ingredientDetailData "alta-cookit-be/features/ingredient_details/data"
)

func ConvertToGorm(entity *ingredients.IngredientEntity) *_ingredientModel.Ingredient {
	gorm := _ingredientModel.Ingredient{
		RecipeID: entity.RecipeID,
		Name:     entity.Name,
		Price:    entity.Price,
		IngredientDetails: *_ingredientDetailData.ConvertToGorms(&entity.IngredientDetailEntities),
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]ingredients.IngredientEntity) *[]_ingredientModel.Ingredient {
	gorms := []_ingredientModel.Ingredient{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_ingredientModel.Ingredient) *ingredients.IngredientEntity {
	return &ingredients.IngredientEntity{
		ID: gorm.ID,
		RecipeID: gorm.RecipeID,
		Name: gorm.Name,
		Price: gorm.Price,
		IngredientDetailEntities: *_ingredientDetailData.ConvertToEntities(&gorm.IngredientDetails),
	}
}

func ConvertToEntities(gorms *[]_ingredientModel.Ingredient) *[]ingredients.IngredientEntity {
	entities := []ingredients.IngredientEntity{}
	for _, gorm := range *gorms {
		entities = append(entities, *ConvertToEntity(&gorm))
	}
	return &entities
}
