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

func ConvertToGorms(entities *[]steps.StepEntity) *[]_stepModel.Step {
	gorms := []_stepModel.Step{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_stepModel.Step) *steps.StepEntity {
	return &steps.StepEntity{
		ID: gorm.ID,
		RecipeID: gorm.RecipeID,
		Name: gorm.Name,
	}
}

func ConvertToEntities(gorms *[]_stepModel.Step) *[]steps.StepEntity {
	entities := []steps.StepEntity{}
	for _, gorm := range *gorms {
		entities = append(entities, *ConvertToEntity(&gorm))
	}
	return &entities
}
