package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID   uint
	RecipeID *uint
}
