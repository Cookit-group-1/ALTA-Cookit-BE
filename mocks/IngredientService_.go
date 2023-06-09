// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	ingredients "alta-cookit-be/features/ingredients"

	mock "github.com/stretchr/testify/mock"
)

// IngredientService_ is an autogenerated mock type for the IngredientService_ type
type IngredientService_ struct {
	mock.Mock
}

// DeleteIngredientById provides a mock function with given fields: ingredientEntity
func (_m *IngredientService_) DeleteIngredientById(ingredientEntity *ingredients.IngredientEntity) error {
	ret := _m.Called(ingredientEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ingredients.IngredientEntity) error); ok {
		r0 = rf(ingredientEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIngredientByRecipeId provides a mock function with given fields: ingredientEntity
func (_m *IngredientService_) DeleteIngredientByRecipeId(ingredientEntity *ingredients.IngredientEntity) error {
	ret := _m.Called(ingredientEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ingredients.IngredientEntity) error); ok {
		r0 = rf(ingredientEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertIngredient provides a mock function with given fields: ingredientEntity
func (_m *IngredientService_) InsertIngredient(ingredientEntity *ingredients.IngredientEntity) (*ingredients.IngredientEntity, error) {
	ret := _m.Called(ingredientEntity)

	var r0 *ingredients.IngredientEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*ingredients.IngredientEntity) (*ingredients.IngredientEntity, error)); ok {
		return rf(ingredientEntity)
	}
	if rf, ok := ret.Get(0).(func(*ingredients.IngredientEntity) *ingredients.IngredientEntity); ok {
		r0 = rf(ingredientEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ingredients.IngredientEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*ingredients.IngredientEntity) error); ok {
		r1 = rf(ingredientEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIngredientById provides a mock function with given fields: ingredientEntity
func (_m *IngredientService_) UpdateIngredientById(ingredientEntity *ingredients.IngredientEntity) error {
	ret := _m.Called(ingredientEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ingredients.IngredientEntity) error); ok {
		r0 = rf(ingredientEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIngredientService_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewIngredientService_ creates a new instance of IngredientService_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIngredientService_(t mockConstructorTestingTNewIngredientService_) *IngredientService_ {
	mock := &IngredientService_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
