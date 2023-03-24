package users

import (
	_userModel "alta-cookit-be/features/users/models"
)

type UserData_ interface {
	SelectUserById(userId uint) *_userModel.User
}
