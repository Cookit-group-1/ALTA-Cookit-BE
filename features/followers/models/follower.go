package models

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	FromUserID uint
	ToUserID uint
}
