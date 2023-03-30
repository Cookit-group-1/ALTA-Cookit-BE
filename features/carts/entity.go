package carts

import "github.com/labstack/echo/v4"

type CartEntity struct {
	ID             uint
	UserID         uint
	IngredientID   uint
	RecipeName     string
	IngredientName string
	Price          float64
	Quantity       int `validate:"required"`
	DataLimit      int
	DataOffset     int
}

type CartRequest struct {
	ID           uint `json:"-" form:"-"`
	UserID       uint `json:"-" form:"-"`
	IngredientID uint `json:"ingredient_id" form:"ingredient_id"`
	Quantity     int  `json:"quantity" form:"quantity"`
	DataLimit    int
	DataOffset   int
}

type CartResponse struct {
	ID             uint    `json:"id,omitempty"`
	RecipeName     string  `json:"recipe_name,omitempty"`
	IngredientName string  `json:"ingredient_name,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
}

type CartDelivery_ interface {
	SelectCartsByUserId(e echo.Context) error
	InsertCart(e echo.Context) error
	UpdateCartById(e echo.Context) error
	DeleteCartById(e echo.Context) error
}

type CartService_ interface {
	SelectCartsByUserId(cartEntity *CartEntity) (*[]CartEntity, error)
	InsertCart(cartEntity *CartEntity) (*CartEntity, error)
	UpdateCartById(cartEntity *CartEntity) error
	DeleteCartById(cartEntity *CartEntity) error
}

type CartData_ interface {
	ActionValidator(id, userId uint) bool
	SelectCartsByUserId(cartEntity *CartEntity) (*[]CartEntity, error)
	InsertCart(cartEntity *CartEntity) (*CartEntity, error)
	UpdateCartById(cartEntity *CartEntity) error
	DeleteCartById(cartEntity *CartEntity) error
}
