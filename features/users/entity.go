package users

import (
	"mime/multipart"

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
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	Deactive(token interface{}) error
}

type UserData interface {
	SelectUserById(existUser Core) *Core
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(userID uint, updateData Core) (Core, error)
	Deactive(userID uint) error
}
