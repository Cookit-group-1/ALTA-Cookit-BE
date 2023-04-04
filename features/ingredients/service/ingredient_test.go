package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertIngredient(t *testing.T) {
	table := InsertIngredientTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)
			mRecipe.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)

			mIngredient := new(mocks.IngredientData_)
			mIngredient.On("InsertIngredient", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mIngredient, mRecipe)
			_, err := service.InsertIngredient(&v.Input.ingredientEntity)
			if index <= 1 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
					assert.Nil(t, err)
				}
			} else if index >= 2 && index <= 3 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 4 && index <= 5 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestUpdateIngredientById(t *testing.T) {
	table := UpdateIngredientByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mIngredient := new(mocks.IngredientData_)
			mIngredient.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mIngredient.On("UpdateIngredientById", mock.Anything).Return(v.Output.errResult)

			service := New(mIngredient, mRecipe)
			err := service.UpdateIngredientById(&v.Input.ingredientEntity)
			if index <= 1 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
					assert.Nil(t, err)
				}
			} else if index >= 2 && index <= 3 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 4 && index <= 5 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestDeleteIngredientStepById(t *testing.T) {
	table := DeleteIngredientByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mIngredient := new(mocks.IngredientData_)
			mIngredient.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mIngredient.On("DeleteIngredientById", mock.Anything).Return(v.Output.errResult)

			service := New(mIngredient, mRecipe)
			err := service.DeleteIngredientById(&v.Input.ingredientEntity)
			if index <= 1 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 2 && index <= 3 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestDeleteIngredientByRecipeId(t *testing.T) {
	table := DeleteIngredientByRecipeIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mIngredient := new(mocks.IngredientData_)
			mIngredient.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mIngredient.On("DeleteIngredientByRecipeId", mock.Anything).Return(v.Output.errResult)

			service := New(mIngredient, mRecipe)
			err := service.DeleteIngredientByRecipeId(&v.Input.ingredientEntity)
			if index <= 1 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 2 && index <= 3 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}
