package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID   uint `gorm:"uniqueIndex:idx_user_recipe"`
	RecipeID uint `gorm:"uniqueIndex:idx_user_recipe"`
}
