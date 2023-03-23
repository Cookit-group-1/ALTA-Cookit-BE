package data

import (
	"alta-cookit-be/features/ingredient_details"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type IngredientDetailData struct {
	db *gorm.DB
}

func New(db *gorm.DB) ingredient_details.IngredientDetailData_ {
	return &IngredientDetailData{
		db: db,
	}
}

func (d *IngredientDetailData) InsertIngredientDetail(entity *ingredient_details.IngredientDetailEntity) (*ingredient_details.IngredientDetailEntity, error) {
	gorm := ConvertToGorm(entity)
	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "ingredient_id") {
			return nil, errors.New(consts.INGREDIENT_InvalidIngredient)
		}
		return nil, tx.Error
	}
	return ConvertToEntity(gorm), nil
}

func (d *IngredientDetailData) UpdateIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	tx := d.db.Where("id = ? AND ingredient_id = ?", entity.ID, entity.IngredientID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *IngredientDetailData) DeleteIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	tx := d.db.Where("id = ? AND ingredient_id = ?", entity.ID, entity.IngredientID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}