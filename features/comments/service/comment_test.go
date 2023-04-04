package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectCommentsByRecipeId(t *testing.T) {
	table := SelectCommentsByRecipeIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mComment := new(mocks.CommentData_)
			mComment.On("SelectCommentsByRecipeId", mock.Anything).Return(v.Output.results, v.Output.errResult)

			service := New(mComment)
			_, err := service.SelectCommentsByRecipeId(&v.Input.commentEntity)
			if index <= 1 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestInsertComment(t *testing.T) {
	table := InsertCommentTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mComment := new(mocks.CommentData_)
			mComment.On("InsertComment", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mComment)
			_, err := service.InsertComment(&v.Input.commentEntity)
			if index <= 1 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
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

func TestUpdateCommentById(t *testing.T) {
	table := UpdateCommentByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mComment := new(mocks.CommentData_)
			mComment.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mComment.On("UpdateCommentById", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mComment)
			_, err := service.UpdateCommentById(&v.Input.commentEntity)
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

func TestDeleteCommentById(t *testing.T) {
	table := DeleteCommentByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mComment := new(mocks.CommentData_)
			mComment.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mComment.On("DeleteCommentById", mock.Anything).Return(v.Output.errResult)

			service := New(mComment)
			err := service.DeleteCommentById(&v.Input.commentEntity)
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