package data

import (
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type IngredientData struct {
	db *gorm.DB
}

func New(db *gorm.DB) ingredients.IngredientData_ {
	return &IngredientData{
		db: db,
	}
}

func (d *IngredientData) InsertIngredient (entity *ingredients.IngredientEntity) (*ingredients.IngredientEntity, error) {
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

func (d *IngredientData) UpdateIngredientById(entity *ingredients.IngredientEntity) error {
	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *IngredientData) DeleteIngredientById(entity *ingredients.IngredientEntity) error {
	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}