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
	return &gorm
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
