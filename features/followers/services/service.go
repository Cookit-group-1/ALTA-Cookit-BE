package services

import (
	"alta-cookit-be/features/followers"
	"errors"
	"strings"
)

type followService struct {
	qry followers.FollowData
}

// Follow implements followers.FollowService
func (fs *followService) Follow(userID uint, followingID uint) error {
	err := fs.qry.Follow(userID, followingID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "already") {
			msg = "you already follow this account"
		} else {
			msg = "data not found"
		}
		return errors.New(msg)
	}
	return nil
}

// ShowAllFollower implements followers.FollowService
func (fs *followService) ShowAllFollower(userID uint) ([]followers.FollowCore, error) {
	res, err := fs.qry.ShowAllFollower(userID)

	if err != nil {
		if strings.Contains(err.Error(), "user") {
			return []followers.FollowCore{}, errors.New("user not found")
		} else {
			return []followers.FollowCore{}, errors.New("internal server error")
		}
	}
	return res, nil
}

// ShowAllFollowing implements followers.FollowService
func (fs *followService) ShowAllFollowing(userID uint) ([]followers.FollowCore, error) {
	res, err := fs.qry.ShowAllFollowing(userID)

	if err != nil {
		if strings.Contains(err.Error(), "user") {
			return []followers.FollowCore{}, errors.New("user not found")
		} else {
			return []followers.FollowCore{}, errors.New("internal server error")
		}
	}
	return res, nil
}

// Unfollow implements followers.FollowService
func (fs *followService) Unfollow(userID uint, followingID uint) error {
	err := fs.qry.Unfollow(userID, followingID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "invalid user id") {
			msg = "invalid user id, data not found"
		}
		return errors.New(msg)
	}
	return nil
}

func New(ud followers.FollowData) followers.FollowService {
	return &followService{
		qry: ud,
	}
}
