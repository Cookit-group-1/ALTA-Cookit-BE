package data

import (
	"alta-cookit-be/features/ingredient_details"
	_ingredientDetailModel "alta-cookit-be/features/ingredient_details/models"
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

func (d *IngredientDetailData) ActionValidator(id, ingredientId, recipeId, userId uint) bool {
	tempGorm := _ingredientDetailModel.IngredientDetail{}
	d.db.Debug().Model(&tempGorm).Joins("left join ingredients igs on igs.id = ingredient_details.ingredient_id").Joins("left join recipes rs on rs.id = igs.recipe_id").Where("ingredient_details.id = ? AND rs.id = ? AND rs.user_id = ?", id, recipeId, userId).Find(&tempGorm)

	return tempGorm.ID != 0
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
	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *IngredientDetailData) DeleteIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}