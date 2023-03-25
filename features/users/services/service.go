package services

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/helpers"
	"errors"
	"log"
	"strings"
)

type userService struct {
	qry users.UserData
}

func New(ud users.UserData) users.UserService {
	return &userService{
		qry: ud,
	}
}

// Login implements users.UserService
func (us *userService) Login(username string, password string) (string, users.Core, error) {
	res, err := us.qry.Login(username)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "empty") {
			msg = "username or password not allowed empty"
		} else {
			msg = "account not registered or server error"
		}
		return "", users.Core{}, errors.New(msg)
	}

	if err := helpers.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", users.Core{}, errors.New("password not matched")
	}

	useToken, _ := middlewares.CreateToken(int(res.ID), res.Role)
	return useToken, res, nil
}

// Register implements users.UserService
func (us *userService) Register(newUser users.Core) (users.Core, error) {
	hashed, err := helpers.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return users.Core{}, errors.New("password process error")
	}

	// err = us.vld.Struct(&newUser)
	// if err != nil {
	// 	log.Println("err", err)
	// 	msg := helpers.ValidationErrorHandle(err)
	// 	return users.Core{}, errors.New(msg)
	// }

	newUser.Password = string(hashed)
	res, err := us.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already used or duplicated"
		} else if strings.Contains(err.Error(), "empty") {
			msg = "username not allowed empty"
		} else {
			msg = "server error"
		}
		return users.Core{}, errors.New(msg)
	}

	return res, nil
}