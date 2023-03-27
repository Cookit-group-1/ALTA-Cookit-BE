package carts

import "github.com/labstack/echo/v4"

type Recipe struct {
	ID         uint
	Name       string
	SellerID   uint
	SellerName uint
}

type Ingredient struct {
	ID       uint
	Quantity uint
	Price    float64
	Recipe   Recipe
}

type Carts struct {
	ID          uint
	Ingredients []Ingredient
	TotalPrice  float64
}

type CartHandler interface {
	AddProduct() echo.HandlerFunc
	ShowAllProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}

type CartService interface {
	AddProduct(userID, ingredientID uint, newCart Carts) (Carts, error)
	ShowAllProduct(userID uint) ([]Carts, error)
	UpdateProduct(userID, ingredientID uint, updateCart Carts) (Carts, error)
	DeleteProduct(userID, ingredientID uint)
}

type CartData interface {
	AddProduct(userID, ingredientID uint, newCart Carts) (Carts, error)
	ShowAllProduct(userID uint) ([]Carts, error)
	UpdateProduct(userID, ingredientID uint, updateCart Carts) (Carts, error)
	DeleteProduct(userID, ingredientID uint)
}
