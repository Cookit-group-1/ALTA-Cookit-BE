// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	carts "alta-cookit-be/features/carts"

	mock "github.com/stretchr/testify/mock"
)

// CartData_ is an autogenerated mock type for the CartData_ type
type CartData_ struct {
	mock.Mock
}

// ActionValidator provides a mock function with given fields: id, userId
func (_m *CartData_) ActionValidator(id uint, userId uint) bool {
	ret := _m.Called(id, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint, uint) bool); ok {
		r0 = rf(id, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteCartById provides a mock function with given fields: cartEntity
func (_m *CartData_) DeleteCartById(cartEntity *carts.CartEntity) error {
	ret := _m.Called(cartEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) error); ok {
		r0 = rf(cartEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertCart provides a mock function with given fields: cartEntity
func (_m *CartData_) InsertCart(cartEntity *carts.CartEntity) (*carts.CartEntity, error) {
	ret := _m.Called(cartEntity)

	var r0 *carts.CartEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) (*carts.CartEntity, error)); ok {
		return rf(cartEntity)
	}
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) *carts.CartEntity); ok {
		r0 = rf(cartEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*carts.CartEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*carts.CartEntity) error); ok {
		r1 = rf(cartEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectCartsByUserId provides a mock function with given fields: cartEntity
func (_m *CartData_) SelectCartsByUserId(cartEntity *carts.CartEntity) (*[]carts.CartEntity, error) {
	ret := _m.Called(cartEntity)

	var r0 *[]carts.CartEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) (*[]carts.CartEntity, error)); ok {
		return rf(cartEntity)
	}
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) *[]carts.CartEntity); ok {
		r0 = rf(cartEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]carts.CartEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*carts.CartEntity) error); ok {
		r1 = rf(cartEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartById provides a mock function with given fields: cartEntity
func (_m *CartData_) UpdateCartById(cartEntity *carts.CartEntity) error {
	ret := _m.Called(cartEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*carts.CartEntity) error); ok {
		r0 = rf(cartEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartData_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData_ creates a new instance of CartData_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData_(t mockConstructorTestingTNewCartData_) *CartData_ {
	mock := &CartData_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
