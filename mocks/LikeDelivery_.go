// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// LikeDelivery_ is an autogenerated mock type for the LikeDelivery_ type
type LikeDelivery_ struct {
	mock.Mock
}

// LikeRecipe provides a mock function with given fields: e
func (_m *LikeDelivery_) LikeRecipe(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnlikeRecipe provides a mock function with given fields: e
func (_m *LikeDelivery_) UnlikeRecipe(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLikeDelivery_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewLikeDelivery_ creates a new instance of LikeDelivery_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLikeDelivery_(t mockConstructorTestingTNewLikeDelivery_) *LikeDelivery_ {
	mock := &LikeDelivery_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
