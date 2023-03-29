package data

import (
	"alta-cookit-be/features/followers"

	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	FromUserID uint `gorm:"uniqueIndex:idx_user"`
	ToUserID   uint `gorm:"uniqueIndex:idx_user"`
}

func DataToCore(data Follower) followers.FollowCore {
	return followers.FollowCore{
		ID:                  data.ID,
		FollowersFromUserID: data.FromUserID,
		FollowingToUserID:   data.ToUserID,
	}
}

func CoreToData(data followers.FollowCore) Follower {
	return Follower{
		Model:      gorm.Model{ID: data.ID},
		FromUserID: data.FollowersFromUserID,
		ToUserID:   data.FollowingToUserID,
	}
}