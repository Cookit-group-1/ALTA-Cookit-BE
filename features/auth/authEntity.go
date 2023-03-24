package auth

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	ProfilePicture string
	Username       string `validate:"required"`
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=3"`
}

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type AuthService interface {
	Register(token interface{}, newUser Core) (Core, error)
	Login(usernam, password string) (string, Core, error)
}
