package steps

import "github.com/labstack/echo/v4"

type StepEntity struct {
	ID       uint
	UserID   uint
	RecipeID uint
	Name     string `validate:"required"`
}

type StepRequest struct {
	ID       uint   `json:"-" form:"-"`
	UserID   uint   `json:"-" form:"-"`
	RecipeID uint   `json:"-" form:"-"`
	Name     string `json:"name" form:"name"`
}

type StepResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type StepDelivery_ interface {
	InsertStep(e echo.Context) error
	UpdateStepById(e echo.Context) error
	DeleteStepById(e echo.Context) error
	DeleteStepByRecipeId(e echo.Context) error
}

type StepService_ interface {
	InsertStep(stepEntity *StepEntity) (*StepEntity, error)
	UpdateStepById(stepEntity *StepEntity) error
	DeleteStepById(stepEntity *StepEntity) error
	DeleteStepByRecipeId(stepEntity *StepEntity) error
}

type StepData_ interface {
	ActionValidator(id, recipeId, userId uint) bool
	InsertStep(stepEntity *StepEntity) (*StepEntity, error)
	UpdateStepById(istepEntity *StepEntity) error
	DeleteStepById(stepEntity *StepEntity) error
	DeleteStepByRecipeId(stepEntity *StepEntity) error
}
