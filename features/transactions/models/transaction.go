package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID       uint
	IngredientID uint
	Type         string  `gorm:"not null;type:enum('Unpaid', 'Paid');default:'Unpaid'"`
	TotalPrice   float64 `gorm:"not null"`
}
