package models

import (
	"gorm.io/gorm"
)

type IngredientDetail struct {
	gorm.Model
	IngredientID uint
	Name         string  `gorm:"not null;type:VARCHAR(50)"`
	Size         float64 `gorm:"not null;type:VARCHAR(50)"`
}
