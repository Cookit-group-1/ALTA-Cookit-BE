package data

import (
	"alta-cookit-be/features/users"
	_userModel "alta-cookit-be/features/users/models"

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

func (d *UserData) SelectUserById(id uint) *_userModel.User {
	tempGorm := _userModel.User{}
	d.db.Where("id = ?", id).Find(&tempGorm)

	if tempGorm.ID == 0 {
		return nil
	}
	return &tempGorm
}
