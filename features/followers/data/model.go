package data

import (
	"alta-cookit-be/features/followers"

	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	FromUserID      uint `gorm:"uniqueIndex:idx_user"`
	ToUserID        uint `gorm:"uniqueIndex:idx_user"`
	AmountFollowing uint
	AmountFollowers uint
}

func DataToCore(data Follower) followers.FollowCore {
	return followers.FollowCore{
		ID:                  data.ID,
		FollowersFromUserID: data.FromUserID,
		FollowingToUserID:   data.ToUserID,
		AmountFollowing:     data.AmountFollowing,
		AmountFollowers:     data.AmountFollowers,
	}
}

func CoreToData(data followers.FollowCore) Follower {
	return Follower{
		Model:           gorm.Model{ID: data.ID},
		FromUserID:      data.FollowersFromUserID,
		ToUserID:        data.FollowingToUserID,
		AmountFollowing: data.AmountFollowing,
		AmountFollowers: data.AmountFollowers,
	}
}
