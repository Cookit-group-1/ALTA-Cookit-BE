package data

import (
	"alta-cookit-be/features/steps"
	_stepModel "alta-cookit-be/features/steps/models"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type StepData struct {
	db *gorm.DB
}

func New(db *gorm.DB) steps.StepData_ {
	return &StepData{
		db: db,
	}
}

func (d *StepData) ActionValidator(id, recipeId, userId uint) bool {
	tempGorm := _stepModel.Step{}
	if id != 0 {
		d.db.Model(&tempGorm).Joins("left join recipes rs on rs.id = steps.recipe_id").Where("steps.id = ? AND rs.id = ? AND rs.user_id = ?", id, recipeId, userId).Find(&tempGorm)
	} else {
		d.db.Model(&tempGorm).Joins("left join recipes rs on rs.id = steps.recipe_id").Where("steps.recipe_id = ? AND rs.user_id = ?", recipeId, userId).First(&tempGorm)
	}

	return tempGorm.ID != 0
}

func (d *StepData) InsertStep (entity *steps.StepEntity) (*steps.StepEntity, error) {
	gorm := ConvertToGorm(entity)
	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		return nil, tx.Error
	}
	return ConvertToEntity(gorm), nil
}

func (d *StepData) UpdateStepById(entity *steps.StepEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *StepData) DeleteStepById(entity *steps.StepEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *StepData) DeleteStepByRecipeId(entity *steps.StepEntity) error {
	tx := d.db.Where("recipe_id = ?", entity.RecipeID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}