package services

import (
	"alta-cookit-be/features/followers"
	"alta-cookit-be/mocks"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
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

func TestShowAllFollowing(t *testing.T) {
	data := mocks.NewFollowData(t)
	filePath := filepath.Join("..", "..", "..", "test.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []followers.FollowCore{{
		ID:             1,
		Username:       "abu",
		Role:           "VerifiedUser",
		ProfilePicture: imageTrueCnv.Filename,
		FromUserID:     9,
		ToUserID:       1,
	}}

	t.Run("success get all following", func(t *testing.T) {
		data.On("ShowAllFollowing", uint(1)).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.ShowAllFollowing(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("ShowAllFollowing", uint(1)).Return([]followers.FollowCore{}, errors.New("internal server error")).Once()
		srv := New(data)
		res, err := srv.ShowAllFollowing(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		data.On("ShowAllFollowing", uint(1)).Return([]followers.FollowCore{}, errors.New("user not found")).Once()
		srv := New(data)
		res, err := srv.ShowAllFollowing(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})
}

func TestShowAllFollower(t *testing.T) {
	data := mocks.NewFollowData(t)
	filePath := filepath.Join("..", "..", "..", "test.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []followers.FollowCore{{
		ID:             1,
		Username:       "abu",
		Role:           "VerifiedUser",
		ProfilePicture: imageTrueCnv.Filename,
		FromUserID:     9,
		ToUserID:       1,
	}}

	t.Run("success get all follower", func(t *testing.T) {
		data.On("ShowAllFollower", uint(1)).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.ShowAllFollower(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("ShowAllFollower", uint(1)).Return([]followers.FollowCore{}, errors.New("internal server error")).Once()
		srv := New(data)
		res, err := srv.ShowAllFollower(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		data.On("ShowAllFollower", uint(1)).Return([]followers.FollowCore{}, errors.New("user not found")).Once()
		srv := New(data)
		res, err := srv.ShowAllFollower(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})
}