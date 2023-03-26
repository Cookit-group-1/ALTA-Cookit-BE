package data

import (
	"alta-cookit-be/features/likes"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type LikeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) likes.LikeData_ {
	return &LikeData{
		db: db,
	}
}

func (d *LikeData) LikeRecipe(entity *likes.LikeEntity) error {
	gorm := ConvertToGorm(entity)

	tx := d.db.Create(&gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return errors.New(consts.RECIPE_InvalidRecipe)
		}
		if strings.Contains(tx.Error.Error(), "idx_user_recipe") {
			return errors.New(consts.LIKE_AlreadyLiked)
		}
		return tx.Error
	}
	return nil
}

func (d *LikeData) UnlikeRecipe(entity *likes.LikeEntity) error {
	gorm := ConvertToGorm(entity)

	tx := d.db.Where("user_id = ? AND recipe_id = ?", entity.UserID, entity.RecipeID).Unscoped().Delete(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
