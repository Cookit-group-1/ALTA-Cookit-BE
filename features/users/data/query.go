package data

import (
	"alta-cookit-be/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData_ {
	return &UserData{
		db: db,
	}
}
