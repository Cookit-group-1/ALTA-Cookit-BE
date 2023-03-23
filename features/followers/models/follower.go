package models

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	FromUserID uint `gorm:"uniqueIndex:idx_user"`
	ToUserID   uint `gorm:"uniqueIndex:idx_user"`
}
