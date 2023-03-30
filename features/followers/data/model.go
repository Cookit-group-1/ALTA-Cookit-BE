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
