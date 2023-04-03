// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	users "alta-cookit-be/features/users"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Deactive provides a mock function with given fields: userID
func (_m *UserService) Deactive(userID uint) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListUserRequest provides a mock function with given fields: userID
func (_m *UserService) ListUserRequest(userID uint) ([]users.Core, error) {
	ret := _m.Called(userID)

	var r0 []users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]users.Core, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []users.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: username, password
func (_m *UserService) Login(username string, password string) (string, users.Core, error) {
	ret := _m.Called(username, password)

	var r0 string
	var r1 users.Core
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, users.Core, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) users.Core); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Get(1).(users.Core)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Profile provides a mock function with given fields: userID
func (_m *UserService) Profile(userID uint) (users.Core, error) {
	ret := _m.Called(userID)

	var r0 users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.Core, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) users.Core); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *UserService) Register(newUser users.Core) (users.Core, error) {
	ret := _m.Called(newUser)

	var r0 users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(users.Core) (users.Core, error)); ok {
		return rf(newUser)
	}
	if rf, ok := ret.Get(0).(func(users.Core) users.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(users.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUser provides a mock function with given fields: userID, quote
func (_m *UserService) SearchUser(userID uint, quote string) ([]users.Core, error) {
	ret := _m.Called(userID, quote)

	var r0 []users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) ([]users.Core, error)); ok {
		return rf(userID, quote)
	}
	if rf, ok := ret.Get(0).(func(uint, string) []users.Core); ok {
		r0 = rf(userID, quote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, quote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, fileData, updateData
func (_m *UserService) Update(userID uint, fileData multipart.FileHeader, updateData users.Core) (users.Core, error) {
	ret := _m.Called(userID, fileData, updateData)

	var r0 users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, multipart.FileHeader, users.Core) (users.Core, error)); ok {
		return rf(userID, fileData, updateData)
	}
	if rf, ok := ret.Get(0).(func(uint, multipart.FileHeader, users.Core) users.Core); ok {
		r0 = rf(userID, fileData, updateData)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(uint, multipart.FileHeader, users.Core) error); ok {
		r1 = rf(userID, fileData, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: userID, updatePassword
func (_m *UserService) UpdatePassword(userID uint, updatePassword users.Core) error {
	ret := _m.Called(userID, updatePassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.Core) error); ok {
		r0 = rf(userID, updatePassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpgradeUser provides a mock function with given fields: userID, approvement
func (_m *UserService) UpgradeUser(userID uint, approvement users.Core) (users.Core, error) {
	ret := _m.Called(userID, approvement)

	var r0 users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, users.Core) (users.Core, error)); ok {
		return rf(userID, approvement)
	}
	if rf, ok := ret.Get(0).(func(uint, users.Core) users.Core); ok {
		r0 = rf(userID, approvement)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(uint, users.Core) error); ok {
		r1 = rf(userID, approvement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
