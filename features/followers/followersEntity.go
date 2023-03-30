package followers

import (
	"github.com/labstack/echo/v4"
)

type FollowCore struct {
	ID              uint
	FromUserID      uint
	ToUserID        uint
	Username        string
	ProfilePicture  string
	AmountFollowing uint
	AmountFollowers uint
	Role            string
	UserID          uint
}

type FollowHandler interface {
	Follow() echo.HandlerFunc
	Unfollow() echo.HandlerFunc
	ShowAllFollower() echo.HandlerFunc
	ShowAllFollowing() echo.HandlerFunc
}

type FollowService interface {
	Follow(userID, followingID uint) error
	Unfollow(userID, followingID uint) error
	ShowAllFollower(userID uint) ([]FollowCore, error)
	ShowAllFollowing(userID uint) ([]FollowCore, error)
}

type FollowData interface {
	Follow(userID, followingID uint) error
	Unfollow(userID, followingID uint) error
	ShowAllFollower(userID uint) ([]FollowCore, error)
	ShowAllFollowing(userID uint) ([]FollowCore, error)
}
