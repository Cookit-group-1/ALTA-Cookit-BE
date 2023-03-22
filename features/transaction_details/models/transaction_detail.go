package models

import (
	"gorm.io/gorm"
)

type TransactionsDetail struct {
	gorm.Model
	TransactionID uint
	IngredientID  uint
	Quantity      int `gorm:"not null;default:1"`
}
