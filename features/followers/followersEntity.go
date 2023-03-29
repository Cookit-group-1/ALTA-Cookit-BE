package followers

import (
	"github.com/labstack/echo/v4"
)

type FollowCore struct {
	ID                  uint
	FollowersFromUserID uint
	FollowingToUserID   uint
	FollowersName       string
	FollowingName       string
	FollowersImage      string
	FollowingImage      string
	AmountFollowing     uint
	AmountFollowers     uint
}

type FollowHandler interface {
	Follow() echo.HandlerFunc
	Unfollow() echo.HandlerFunc
	ShowAllFollower() echo.HandlerFunc
	ShowAllFollowing() echo.HandlerFunc
}

type FollowService interface {
	Follow(userID uint) error
	Unfollow(userID uint) error
	ShowAllFollower() ([]FollowCore, error)
	ShowAllFollowing() ([]FollowCore, error)
}

type FollowData interface {
	Follow(userID uint) error
	Unfollow(userID uint) error
	ShowAllFollower() ([]FollowCore, error)
	ShowAllFollowing() ([]FollowCore, error)
}
