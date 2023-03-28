package models

import (
	_cartModel "alta-cookit-be/features/carts/models"
	_ingredientDetailModel "alta-cookit-be/features/ingredient_details/models"
	_transactionModel "alta-cookit-be/features/transactions/models"

	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	RecipeID          uint
	Name              string                                    `gorm:"default:'';not null;type:VARCHAR(255)"`
	Price             float64                                   `gorm:"not null;default:0"`
	IngredientDetails []_ingredientDetailModel.IngredientDetail `gorm:"constraint:OnDelete:CASCADE;"`
	Carts             []_cartModel.Cart                         `gorm:"constraint:OnDelete:CASCADE;"`
	Transactions      []_transactionModel.Transaction           `gorm:"constraint:OnDelete:CASCADE;"`
	IngredientRefer uint
}
