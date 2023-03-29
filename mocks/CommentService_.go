// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	comments "alta-cookit-be/features/comments"

	mock "github.com/stretchr/testify/mock"
)

// CommentService_ is an autogenerated mock type for the CommentService_ type
type CommentService_ struct {
	mock.Mock
}

// DeleteCommentById provides a mock function with given fields: commentEntity
func (_m *CommentService_) DeleteCommentById(commentEntity *comments.CommentEntity) error {
	ret := _m.Called(commentEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) error); ok {
		r0 = rf(commentEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertComment provides a mock function with given fields: commentEntity
func (_m *CommentService_) InsertComment(commentEntity *comments.CommentEntity) (*comments.CommentEntity, error) {
	ret := _m.Called(commentEntity)

	var r0 *comments.CommentEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) (*comments.CommentEntity, error)); ok {
		return rf(commentEntity)
	}
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) *comments.CommentEntity); ok {
		r0 = rf(commentEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comments.CommentEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*comments.CommentEntity) error); ok {
		r1 = rf(commentEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectCommentsByRecipeId provides a mock function with given fields: commentEntity
func (_m *CommentService_) SelectCommentsByRecipeId(commentEntity *comments.CommentEntity) (*[]comments.CommentEntity, error) {
	ret := _m.Called(commentEntity)

	var r0 *[]comments.CommentEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) (*[]comments.CommentEntity, error)); ok {
		return rf(commentEntity)
	}
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) *[]comments.CommentEntity); ok {
		r0 = rf(commentEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]comments.CommentEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*comments.CommentEntity) error); ok {
		r1 = rf(commentEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCommentById provides a mock function with given fields: commentEntity
func (_m *CommentService_) UpdateCommentById(commentEntity *comments.CommentEntity) (*comments.CommentEntity, error) {
	ret := _m.Called(commentEntity)

	var r0 *comments.CommentEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) (*comments.CommentEntity, error)); ok {
		return rf(commentEntity)
	}
	if rf, ok := ret.Get(0).(func(*comments.CommentEntity) *comments.CommentEntity); ok {
		r0 = rf(commentEntity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comments.CommentEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(*comments.CommentEntity) error); ok {
		r1 = rf(commentEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentService_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentService_ creates a new instance of CommentService_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentService_(t mockConstructorTestingTNewCommentService_) *CommentService_ {
	mock := &CommentService_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
