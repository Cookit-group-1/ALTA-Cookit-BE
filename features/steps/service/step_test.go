package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertStep(t *testing.T) {
	table := InsertStepTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)
			mRecipe.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)

			mStep := new(mocks.StepData_)
			mStep.On("InsertStep", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mStep, mRecipe)
			_, err := service.InsertStep(&v.Input.stepEntity)
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

func TestUpdateStepById(t *testing.T) {
	table := UpdateStepByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mStep := new(mocks.StepData_)
			mStep.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mStep.On("UpdateStepById", mock.Anything).Return(v.Output.errResult)

			service := New(mStep, mRecipe)
			err := service.UpdateStepById(&v.Input.stepEntity)
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

func TestDeleteStepById(t *testing.T) {
	table := DeleteStepByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mStep := new(mocks.StepData_)
			mStep.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mStep.On("DeleteStepById", mock.Anything).Return(v.Output.errResult)

			service := New(mStep, mRecipe)
			err := service.DeleteStepById(&v.Input.stepEntity)
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

func TestDeleteStepByRecipeId(t *testing.T) {
	table := DeleteStepByRecipeIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mStep := new(mocks.StepData_)
			mStep.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mStep.On("DeleteStepByRecipeId", mock.Anything).Return(v.Output.errResult)

			service := New(mStep, mRecipe)
			err := service.DeleteStepByRecipeId(&v.Input.stepEntity)
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
