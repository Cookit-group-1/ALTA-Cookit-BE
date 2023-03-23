package ingredients

import "github.com/labstack/echo/v4"

type IngredientEntity struct {
	ID       uint
	UserID   uint
	UserRole uint
	RecipeID uint
	Name     string
	Price    float64
}

type IngredientRequest struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	UserRole uint    `json:"user_role"`
	RecipeID uint    `json:"recipe_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

type IngredientResponse struct {
	ID    uint    `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type IngredientDelivery_ interface {
	InsertIngredient(e echo.Context) error
	UpdateIngredientById(e echo.Context) error
	DeleteIngredientById(e echo.Context) error
}

type IngredientService_ interface {
	InsertIngredient(ingredientEntity *IngredientEntity) (*IngredientEntity, error)
	UpdateIngredientById(ingredientEntity *IngredientEntity) error
	DeleteIngredientById(ingredientEntity *IngredientEntity) error
}

type IngredientData_ interface {
	InsertIngredient(ingredientEntity *IngredientEntity) (*IngredientEntity, error)
	UpdateIngredientById(ingredientEntity *IngredientEntity) error
	DeleteIngredientById(ingredientEntity *IngredientEntity) error
}
