package data

import (
	"alta-cookit-be/features/carts"

	"gorm.io/gorm"
)

type CartQuery struct {
	db *gorm.DB
}

// AddProduct implements carts.CartData
func (cq *CartQuery) AddProduct(userID uint, ingredientID uint, newCart carts.CartsCore) (carts.CartsCore, error) {
	panic("unimplemented")
}

// DeleteProduct implements carts.CartData
func (*CartQuery) DeleteProduct(userID uint, ingredientID uint) {
	panic("unimplemented")
}

// ShowAllProduct implements carts.CartData
func (*CartQuery) ShowAllProduct(userID uint) ([]carts.CartsCore, error) {
	panic("unimplemented")
}

// UpdateProduct implements carts.CartData
func (*CartQuery) UpdateProduct(userID uint, ingredientID uint, updateCart carts.CartsCore) (carts.CartsCore, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) carts.CartData {
	return &CartQuery{
		db: db,
	}
}
