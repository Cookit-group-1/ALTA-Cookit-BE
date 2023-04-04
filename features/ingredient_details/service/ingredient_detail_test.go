package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertIngredientDetail(t *testing.T) {
	table := InsertIngredientDetailTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)
			mRecipe.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)

			mIngredientDetail := new(mocks.IngredientDetailData_)
			mIngredientDetail.On("InsertIngredientDetail", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mIngredientDetail, mRecipe)
			_, err := service.InsertIngredientDetail(&v.Input.ingredientDetailEntity)
			if index <= 2 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
					assert.Nil(t, err)
				}
			} else if index >= 3 && index <= 4 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 5 && index <= 6 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestUpdateIngredientDetailById(t *testing.T) {
	table := UpdateIngredientDetailByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mIngredientDetail := new(mocks.IngredientDetailData_)
			mIngredientDetail.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mIngredientDetail.On("UpdateIngredientDetailById", mock.Anything).Return(v.Output.errResult)

			service := New(mIngredientDetail, mRecipe)
			err := service.UpdateIngredientDetailById(&v.Input.ingredientDetailEntity)
			if index <= 2 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
					assert.Nil(t, err)
				}
			} else if index >= 3 && index <= 4 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 5 && index <= 6 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestDeleteIngredientDetailById(t *testing.T) {
	table := DeleteIngredientDetailByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mIngredientDetail := new(mocks.IngredientDetailData_)
			mIngredientDetail.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mIngredientDetail.On("DeleteIngredientDetailById", mock.Anything).Return(v.Output.errResult)

			service := New(mIngredientDetail, mRecipe)
			err := service.DeleteIngredientDetailById(&v.Input.ingredientDetailEntity)
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