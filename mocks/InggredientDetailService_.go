// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	ingredient_details "alta-cookit-be/features/ingredient_details"

	mock "github.com/stretchr/testify/mock"
)

// IngredientDetailService_ is an autogenerated mock type for the IngredientDetailService_ type
type IngredientDetailService_ struct {
	mock.Mock
}

// DeleteIngredientDetailById provides a mock function with given fields: ingredientDetailEntity
func (_m *IngredientDetailService_) DeleteIngredientDetailById(ingredientDetailEntity *ingredient_details.IngredientDetailEntity) error {
	ret := _m.Called(ingredientDetailEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ingredient_details.IngredientDetailEntity) error); ok {
		r0 = rf(ingredientDetailEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertIngredientDetail provides a mock function with given fields: ingredientDetailEntity
func (_m *IngredientDetailService_) InsertIngredientDetail(ingredientDetailEntity *ingredient_details.IngredientDetailEntity) (*ingredient_details.IngredientDetailEntity, error) {
	ret := _m.Called(ingredientDetailEntity)

	var r0 *ingredient_details.IngredientDetailEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*ingredient_details.IngredientDetailEntity) (*ingredient_details.IngredientDetailEntity, error)); ok {
		return rf(ingredientDetailEntity)
	}
	if rf, ok := ret.Get(0).(func(*ingredient_details.IngredientDetailEntity) *ingredient_details.IngredientDetailEntity); ok {
		r0 = rf(ingredientDetailEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ingredient_details.IngredientDetailEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*ingredient_details.IngredientDetailEntity) error); ok {
		r1 = rf(ingredientDetailEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIngredientDetailById provides a mock function with given fields: ingredientDetailEntity
func (_m *IngredientDetailService_) UpdateIngredientDetailById(ingredientDetailEntity *ingredient_details.IngredientDetailEntity) error {
	ret := _m.Called(ingredientDetailEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ingredient_details.IngredientDetailEntity) error); ok {
		r0 = rf(ingredientDetailEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIngredientDetailService_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewIngredientDetailService_ creates a new instance of IngredientDetailService_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIngredientDetailService_(t mockConstructorTestingTNewIngredientDetailService_) *IngredientDetailService_ {
	mock := &IngredientDetailService_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}