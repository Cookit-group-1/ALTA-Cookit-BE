// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	followers "alta-cookit-be/features/followers"

	mock "github.com/stretchr/testify/mock"
)

// FollowData is an autogenerated mock type for the FollowData type
type FollowData struct {
	mock.Mock
}

// Follow provides a mock function with given fields: userID, followingID
func (_m *FollowData) Follow(userID uint, followingID uint) error {
	ret := _m.Called(userID, followingID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, followingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShowAllFollower provides a mock function with given fields: userID
func (_m *FollowData) ShowAllFollower(userID uint) ([]followers.FollowCore, error) {
	ret := _m.Called(userID)

	var r0 []followers.FollowCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]followers.FollowCore, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []followers.FollowCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]followers.FollowCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAllFollowing provides a mock function with given fields: userID
func (_m *FollowData) ShowAllFollowing(userID uint) ([]followers.FollowCore, error) {
	ret := _m.Called(userID)

	var r0 []followers.FollowCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]followers.FollowCore, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []followers.FollowCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]followers.FollowCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unfollow provides a mock function with given fields: userID, followingID
func (_m *FollowData) Unfollow(userID uint, followingID uint) error {
	ret := _m.Called(userID, followingID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, followingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFollowData interface {
	mock.TestingT
	Cleanup(func())
}

// NewFollowData creates a new instance of FollowData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFollowData(t mockConstructorTestingTNewFollowData) *FollowData {
	mock := &FollowData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
