package data

import (
	"alta-cookit-be/features/followers"

	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	FromUserID     uint `gorm:"uniqueIndex:idx_user"`
	ToUserID       uint `gorm:"uniqueIndex:idx_user"`
	Username       string
	ProfilePicture string
	Role           string
	UserID         uint
	UserRefer      uint
}

type Followings struct {
	ID             uint
	Username       string
	ProfilePicture string
	Role           string
	ToUserID       uint
}

func FollowingDataToCore(data Followings) followers.FollowCore {
	return followers.FollowCore{
		UserID:             data.ID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
		ToUserID:       data.ToUserID,
	}
}

func CoreToFollowingData(data followers.FollowCore) Followings {
	return Followings{
		ID:             data.UserID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
		ToUserID:       data.ToUserID,
	}
}

func DataToCore(data Follower) followers.FollowCore {
	return followers.FollowCore{
		ID:             data.ID,
		FromUserID:     data.FromUserID,
		ToUserID:       data.ToUserID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
		UserID:         data.UserID,
	}
}

func CoreToData(data followers.FollowCore) Follower {
	return Follower{
		Model:          gorm.Model{ID: data.ID},
		FromUserID:     data.FromUserID,
		ToUserID:       data.ToUserID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
		UserID:         data.UserID,
	}
}
