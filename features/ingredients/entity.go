package ingredients

import (
	"alta-cookit-be/features/ingredient_details"

	"github.com/labstack/echo/v4"
)

type IngredientEntity struct {
	ID                       uint
	UserID                   uint
	UserRole                 string
	RecipeID                 uint
	Name                     string `validate:"required"`
	Price                    float64
	IngredientDetailEntities []ingredient_details.IngredientDetailEntity
}

type IngredientRequest struct {
	ID                       uint                                         `json:"-" form:"-"`
	UserID                   uint                                         `json:"-" form:"-"`
	UserRole                 string                                       `json:"-" form:"-"`
	RecipeID                 uint                                         `json:"-" form:"-"`
	Name                     string                                       `json:"name" form:"name"`
	Price                    float64                                      `json:"price" form:"price"`
	IngredientDetailRequests []ingredient_details.IngredientDetailRequest `json:"ingredient_details" form:"ingredient_details"`
}

type IngredientResponse struct {
	ID                        uint                                          `json:"id,omitempty"`
	Name                      string                                        `json:"name,omitempty"`
	Price                     float64                                       `json:"price,omitempty"`
	IngredientDetailResponses []ingredient_details.IngredientDetailResponse `json:"ingredient_details,omitempty"`
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
	ActionValidator(id, recipeId, userId uint) bool
	InsertIngredient(ingredientEntity *IngredientEntity) (*IngredientEntity, error)
	UpdateIngredientById(ingredientEntity *IngredientEntity) error
	DeleteIngredientById(ingredientEntity *IngredientEntity) error
}
