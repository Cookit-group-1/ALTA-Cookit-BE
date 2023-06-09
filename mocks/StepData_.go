// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	steps "alta-cookit-be/features/steps"

	mock "github.com/stretchr/testify/mock"
)

// StepData_ is an autogenerated mock type for the StepData_ type
type StepData_ struct {
	mock.Mock
}

// ActionValidator provides a mock function with given fields: id, recipeId, userId
func (_m *StepData_) ActionValidator(id uint, recipeId uint, userId uint) bool {
	ret := _m.Called(id, recipeId, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint, uint, uint) bool); ok {
		r0 = rf(id, recipeId, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteStepById provides a mock function with given fields: stepEntity
func (_m *StepData_) DeleteStepById(stepEntity *steps.StepEntity) error {
	ret := _m.Called(stepEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*steps.StepEntity) error); ok {
		r0 = rf(stepEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStepByRecipeId provides a mock function with given fields: stepEntity
func (_m *StepData_) DeleteStepByRecipeId(stepEntity *steps.StepEntity) error {
	ret := _m.Called(stepEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*steps.StepEntity) error); ok {
		r0 = rf(stepEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertStep provides a mock function with given fields: stepEntity
func (_m *StepData_) InsertStep(stepEntity *steps.StepEntity) (*steps.StepEntity, error) {
	ret := _m.Called(stepEntity)

	var r0 *steps.StepEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*steps.StepEntity) (*steps.StepEntity, error)); ok {
		return rf(stepEntity)
	}
	if rf, ok := ret.Get(0).(func(*steps.StepEntity) *steps.StepEntity); ok {
		r0 = rf(stepEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*steps.StepEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*steps.StepEntity) error); ok {
		r1 = rf(stepEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStepById provides a mock function with given fields: istepEntity
func (_m *StepData_) UpdateStepById(istepEntity *steps.StepEntity) error {
	ret := _m.Called(istepEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*steps.StepEntity) error); ok {
		r0 = rf(istepEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStepData_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewStepData_ creates a new instance of StepData_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStepData_(t mockConstructorTestingTNewStepData_) *StepData_ {
	mock := &StepData_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
