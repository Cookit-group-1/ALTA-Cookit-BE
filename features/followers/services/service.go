package services

import (
	"alta-cookit-be/features/followers"
	"errors"
	"log"
)

type followService struct {
	qry followers.FollowData
}

// Follow implements followers.FollowService
func (fs *followService) Follow(userID uint, followingID uint) error {
	err := fs.qry.Follow(userID, followingID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, following account fail")
	}
	return nil
}

// ShowAllFollower implements followers.FollowService
func (fs *followService) ShowAllFollower() ([]followers.FollowCore, error) {
	panic("unimplemented")
}

// ShowAllFollowing implements followers.FollowService
func (fs *followService) ShowAllFollowing() ([]followers.FollowCore, error) {
	panic("unimplemented")
}

// Unfollow implements followers.FollowService
func (fs *followService) Unfollow(userID uint, followingID uint) error {
	panic("unimplemented")
}

func New(ud followers.FollowData) followers.FollowService {
	return &followService{
		qry: ud,
	}
}
