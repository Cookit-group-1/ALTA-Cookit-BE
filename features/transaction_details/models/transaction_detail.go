package models

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	IngredientID  uint
	Quantity      int `gorm:"not null;default:1"`
}
