package users

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	ProfilePicture string
	Username       string `validate:"required"`
	Bio            string
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=5"`
	Role           string `validate:"required"`
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
}
