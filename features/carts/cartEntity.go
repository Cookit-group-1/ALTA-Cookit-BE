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
	AddRecipe() echo.HandlerFunc
	ShowAllRecipe() echo.HandlerFunc
	UpdateRecipe() echo.HandlerFunc
	DeleteRecipe() echo.HandlerFunc
}
