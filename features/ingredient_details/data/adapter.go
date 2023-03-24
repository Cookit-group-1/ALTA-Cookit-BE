package data

import (
	"alta-cookit-be/features/ingredient_details"
	_ingredientDetailModel "alta-cookit-be/features/ingredient_details/models"
)

func ConvertToGorm(entity *ingredient_details.IngredientDetailEntity) *_ingredientDetailModel.IngredientDetail {
	gorm := _ingredientDetailModel.IngredientDetail{
		IngredientID: entity.IngredientID,
		Name:         entity.Name,
		Quantity:     entity.Quantity,
		Unit:         entity.Unit,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]ingredient_details.IngredientDetailEntity) *[]_ingredientDetailModel.IngredientDetail {
	gorms := []_ingredientDetailModel.IngredientDetail{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_ingredientDetailModel.IngredientDetail) *ingredient_details.IngredientDetailEntity {
	return &ingredient_details.IngredientDetailEntity{
		ID: gorm.ID,
		IngredientID: gorm.IngredientID,
		Name: gorm.Name,
		Quantity: gorm.Quantity,
		Unit: gorm.Unit,
	}
}

func ConvertToEntities(gorms *[]_ingredientDetailModel.IngredientDetail) *[]ingredient_details.IngredientDetailEntity {
	entities := []ingredient_details.IngredientDetailEntity{}
	for _, gorm := range *gorms {
		entities = append(entities, *ConvertToEntity(&gorm))
	}
	return &entities
}
