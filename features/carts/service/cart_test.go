package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectCartsByUserId(t *testing.T) {
	table := SelectCartsByUserIdTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mCart := new(mocks.CartData_)
			mCart.On("SelectCartsByUserId", mock.Anything).Return(v.Output.results, v.Output.errResult)

			service := New(mCart)
			_, err := service.SelectCartsByUserId(&v.Input.cartEntity)
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

func TestInsertCart(t *testing.T) {
	table := InsertCartTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mCart := new(mocks.CartData_)
			mCart.On("InsertCart", mock.Anything).Return(v.Output.results, v.Output.errResult)

			service := New(mCart)
			_, err := service.InsertCart(&v.Input.cartEntity)
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

func TestUpdateCartById(t *testing.T) {
	table := UpdateCartByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mCart := new(mocks.CartData_)
			mCart.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mCart.On("UpdateCartById", mock.Anything).Return(v.Output.errResult)

			service := New(mCart)
			err := service.UpdateCartById(&v.Input.cartEntity)
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

func TestDeleteCartById(t *testing.T) {
	table := DeleteCartByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mCart := new(mocks.CartData_)
			mCart.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mCart.On("DeleteCartById", mock.Anything).Return(v.Output.errResult)

			service := New(mCart)
			err := service.DeleteCartById(&v.Input.cartEntity)
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