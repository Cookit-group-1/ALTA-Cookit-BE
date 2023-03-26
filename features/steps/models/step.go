package models

import (
	"gorm.io/gorm"
)

type Step struct {
	gorm.Model
	RecipeID    uint
	Name        string `gorm:"default:'';not null;type:text"`
}
