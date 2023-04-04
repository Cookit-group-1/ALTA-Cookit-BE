package services

import (
	"alta-cookit-be/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFollow(t *testing.T) {
	data := mocks.NewFollowData(t)
	srv := New(data)

	t.Run("success following account", func(t *testing.T) {
		data.On("Follow", uint(1), mock.Anything).Return(nil).Once()
		err := srv.Follow(uint(1), uint(2))
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("you already follow this account", func(t *testing.T) {
		data.On("Follow", uint(1), mock.Anything).Return(errors.New("you already follow this account")).Once()
		err := srv.Follow(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "already")
		data.AssertExpectations(t)
	})
	t.Run("data not found", func(t *testing.T) {
		data.On("Follow", uint(1), mock.Anything).Return(errors.New("data not found")).Once()
		err := srv.Follow(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

}

func TestUnfollow(t *testing.T) {
	data := mocks.NewFollowData(t)
	srv := New(data)

	t.Run("success unfollowing account", func(t *testing.T) {
		data.On("Unfollow", uint(1), mock.Anything).Return(nil).Once()
		err := srv.Unfollow(uint(1), uint(2))
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("invalid user id, data not found", func(t *testing.T) {
		data.On("Unfollow", uint(1), mock.Anything).Return(errors.New("invalid user id, data not found")).Once()
		err := srv.Unfollow(uint(1), uint(2))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "invalid user id")
		data.AssertExpectations(t)
	})

}
