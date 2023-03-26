package helpers

import (
	"alta-cookit-be/features/users"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator"
)

func TypeFile(test multipart.File) (string, error) {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	TipenamaFile := ""
	if fileType == "image/png" {
		TipenamaFile = ".png"
	} else {
		TipenamaFile = ".jpg"
	}
	if fileType == "image/png" || fileType == "image/jpeg" || fileType == "image/jpg" {
		return TipenamaFile, nil
	}
	return "", errors.New("file type not match")
}

type UserValidate struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3,alphanum"`
}

func CoreToRegVal(data users.Core) UserValidate {
	return UserValidate{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}
func RegistrationValidate(data users.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
			if e.Value() == "" {
				vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(vlderror)
			} else {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
		}
	}
	return nil
}
