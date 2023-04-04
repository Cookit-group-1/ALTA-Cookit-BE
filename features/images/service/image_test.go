package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertImage(t *testing.T) {
	table := InsertImageTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)
			mRecipe.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)

			mImage := new(mocks.ImageData_)
			mImage.On("InsertImage", mock.Anything).Return(v.Output.results, v.Output.errResult)

			service := New(mImage, mRecipe)
			_, err := service.InsertImage(v.Input.imageEntities)
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

func TestUpdateImageById(t *testing.T) {
	table := UpdateImageByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mImage := new(mocks.ImageData_)
			mImage.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mImage.On("UpdateImageById", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mImage, mRecipe)
			_, err := service.UpdateImageById(v.Input.imageEntity)
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

func TestDeleteImageById(t *testing.T) {
	table := DeleteImageByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mImage := new(mocks.ImageData_)
			mImage.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mImage.On("DeleteImageById", mock.Anything).Return(v.Output.errResult)

			service := New(mImage, mRecipe)
			err := service.DeleteImageById(v.Input.imageEntity)
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

func TestDeleteImageByRecipeId(t *testing.T) {
	table := DeleteImageByRecipeIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mRecipe := new(mocks.RecipeData_)

			mImage := new(mocks.ImageData_)
			mImage.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mImage.On("DeleteImageByRecipeId", mock.Anything).Return(v.Output.errResult)

			service := New(mImage, mRecipe)
			err := service.DeleteImageByRecipeId(v.Input.imageEntity)
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
