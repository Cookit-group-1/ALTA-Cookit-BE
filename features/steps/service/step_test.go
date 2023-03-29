package service

import (
	// "alta-cookit-be/mocks"
	// "testing"

	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

// func TestLikeRecipe(t *testing.T) {
// 	table := LikeRecipeTestTable()
// 	for _, v := range table {
// 		t.Run(v.Name, func(t *testing.T) {
// 			m := new(mocks.LikeData_)
// 			m.On("LikeRecipe", mock.Anything).Return(v.Output.errResult)

// 			service := New(m)
// 			err := service.LikeRecipe(&v.Input.likeEntity)
// 			if v.Output.IsError {
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

// func TestUnlikeRecipe(t *testing.T) {
// 	table := LikeRecipeTestTable()
// 	for _, v := range table {
// 		t.Run(v.Name, func(t *testing.T) {
// 			m := new(mocks.LikeData_)
// 			m.On("UnlikeRecipe", mock.Anything).Return(v.Output.errResult)

// 			service := New(m)
// 			err := service.UnlikeRecipe(&v.Input.likeEntity)
// 			if v.Output.IsError {
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }
