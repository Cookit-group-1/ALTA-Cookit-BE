package users

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                   uint
	ProfilePicture       string
	Username             string
	Bio                  string
	Email                string
	Password             string
	NewPassword          string
	ConfirmationPassword string
	Role                 string
	Approvement          string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactive() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	UpgradeUser() echo.HandlerFunc
	SearchUser() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
	Profile(userID uint) (Core, error)
	Update(userID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
	Deactive(userID uint) error
	UpdatePassword(userID uint, updatePassword Core) error
	UpgradeUser(userID uint, approvement Core) (Core, error)
	SearchUser(quote string) ([]Core, error)
}

type UserData interface {
	SelectUserById(existUser Core) *Core
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(userID uint, updateData Core) (Core, error)
	Deactive(userID uint) error
	UpgradeUser(userID uint, approvement Core) (Core, error)
	SearchUser(quote string) ([]Core, error)
}
