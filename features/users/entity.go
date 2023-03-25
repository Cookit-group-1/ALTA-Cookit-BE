package users

import (
	_userModel "alta-cookit-be/features/users/models"
	"github.com/labstack/echo/v4"
)

type UserData_ interface {
	SelectUserById(userId uint) *_userModel.User
}

type Core struct {
	ID             uint
	ProfilePicture string
	Username       string `validate:"required"`
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=5"`
	Role           string `validate:"required"`
}

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type AuthService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
}

type AuthData interface {
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
}