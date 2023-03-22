package models

import (
	"gorm.io/gorm"
)

type Step struct {
	gorm.Model
	RecipeID    uint
	Name        string `gorm:"not null;type:text"`
}
