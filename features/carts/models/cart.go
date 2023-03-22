package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID       uint
	IngredientID uint
	Quantity     int `gorm:"not null;default:1"`
}
