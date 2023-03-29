package models

import (
	"alta-cookit-be/features/carts"

	"gorm.io/gorm"
	// ingredient "alta-cookit-be/features/ingredients/models"
)

type Cart struct {
	gorm.Model
	UserID       uint
	IngredientID uint
	Quantity     uint `gorm:"not null;default:1"`
	TotalPrice   float64
	// Ingredient   []Ingredient `gorm:"foreignKey:IngredientRefer"`
}

// type Ingredient struct {
// 	gorm.Model
// 	Quantity        uint
// 	Price           float64
// 	// IngredientRefer uint
// 	Recipe          Recipe `gorm:"foreignKey:RecipeName"`
// }

// type Recipe struct {
// 	gorm.Model
// 	RecipeName  string
// 	RecipeImage string
// 	SellerID    uint
// 	SellerName  string
// }

func CartToCore(data Cart) carts.CartsCore {
	return carts.CartsCore{
		ID:          data.ID,
		TotalPrice:  data.TotalPrice,
		Quantity:    data.Quantity,
		Ingredients: []carts.IngredientCore{},
	}
}
