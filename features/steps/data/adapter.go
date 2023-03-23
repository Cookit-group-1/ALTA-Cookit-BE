package data

import (
	_stepModel "alta-cookit-be/features/steps/models"
	"alta-cookit-be/features/steps"
)

func ConvertToGorm(entity *steps.StepEntity) *_stepModel.Step {
	gorm := _stepModel.Step{
		RecipeID: entity.RecipeID,
		Name:     entity.Name,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToEntity(gorm *_stepModel.Step) *steps.StepEntity {
	return &steps.StepEntity{
		ID: gorm.ID,
		RecipeID: gorm.RecipeID,
		Name: gorm.Name,
	}
}
