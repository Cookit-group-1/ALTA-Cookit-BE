package helpers

import (
	"alta-cookit-be/features/users"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

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

type UserRegisterValidate struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,secure_password"`
}

// func CoreUserReqToVal(data users.Core) UserRegisterValidate {
// 	return UserRegisterValidate{
// 		Username: data.Username,
// 		Email:    data.Email,
// 		Password: data.Password,
// 	}
// }

type UserLoginValidate struct {
	Username string `validate:"required"`
	Password string `validate:"required,secure_password"`
}

type PasswordValidate struct {
	Password string `validate:"secure_password"`
}

type EmailValidate struct {
	Email string `validate:"email"`
}

func ToValidate(option string, data interface{}) interface{} {
	switch option {
	case "register":
		res := UserRegisterValidate{}
		if v, ok := data.(users.Core); ok {
			res.Username = v.Username
			res.Email = v.Email
			res.Password = v.Password
		}
		return res
	case "login":
		res := UserLoginValidate{}
		if v, ok := data.(users.Core); ok {
			res.Username = v.Username
			res.Password = v.Password
		}
		return res
	case "password":
		res := PasswordValidate{}
		if v, ok := data.(users.Core); ok {
			res.Password = v.Password
		}
		return res
	case "email":
		res := EmailValidate{}
		if v, ok := data.(users.Core); ok {
			res.Email = v.Email
		}
		return res
	default:
		return nil
	}
}

func securePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	if regexp.MustCompile(`^(?i)(password|1234|qwerty)`).MatchString(password) {
		return false
	}
	return true
}

func Validation(data interface{}) error {
	validate = validator.New()
	validate.RegisterValidation("secure_password", securePassword)
	err := validate.Struct(data)
	if err != nil {
		log.Println("log on helper validation: ", err)
		return err
	}
	return nil
}
