package ingredient_details

import "github.com/labstack/echo/v4"

type IngredientDetailEntity struct {
	ID           uint
	UserID       uint
	UserRole     string
	IngredientID uint
	Name         string
	Quantity     int
	Unit         string
}

type IngredientDetailRequest struct {
	ID           uint `json:"id"`
	UserID       uint
	UserRole     string `json:"user_role"`
	IngredientID uint
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	Unit         string `json:"unit"`
}

type IngredientDetailResponse struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

type IngredientDetailDelivery_ interface {
	InsertIngredientDetail(e echo.Context) error
	UpdateIngredientDetailById(e echo.Context) error
	DeleteIngredientDetailById(e echo.Context) error
}

type IngredientDetailService_ interface {
	InsertIngredientDetail(ingredientDetailEntity *IngredientDetailEntity) (*IngredientDetailEntity, error)
	UpdateIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
	DeleteIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
}

type IngredientDetailData_ interface {
	InsertIngredientDetail(ingredientDetailEntity *IngredientDetailEntity) (*IngredientDetailEntity, error)
	UpdateIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
	DeleteIngredientDetailById(ingredientDetailEntity *IngredientDetailEntity) error
}
