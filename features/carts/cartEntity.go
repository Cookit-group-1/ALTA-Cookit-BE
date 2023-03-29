package carts

import "github.com/labstack/echo/v4"

type RecipeCore struct {
	ID          uint
	Name        string
	SellerID    uint
	SellerName  uint
	RecipeImage string
}

type IngredientCore struct {
	ID       uint
	Quantity uint
	Price    float64
	Recipe   RecipeCore
}

type CartsCore struct {
	ID          uint
	Ingredients []IngredientCore
	Quantity    uint
	TotalPrice  float64
}

type CartHandler interface {
	AddProduct() echo.HandlerFunc
	ShowAllProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}

type CartService interface {
	AddProduct(userID, ingredientID uint, newCart CartsCore) (CartsCore, error)
	ShowAllProduct(userID uint) ([]CartsCore, error)
	UpdateProduct(userID, ingredientID uint, updateCart CartsCore) (CartsCore, error)
	DeleteProduct(userID, ingredientID uint)
}

type CartData interface {
	AddProduct(userID, ingredientID uint, newCart CartsCore) (CartsCore, error)
	ShowAllProduct(userID uint) ([]CartsCore, error)
	UpdateProduct(userID, ingredientID uint, updateCart CartsCore) (CartsCore, error)
	DeleteProduct(userID, ingredientID uint)
}
