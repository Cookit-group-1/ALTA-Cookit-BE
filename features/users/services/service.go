package services

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/helpers"
	"errors"
	"log"
	"mime/multipart"
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

// Deactive implements users.UserService
func (us *userService) Deactive(userID uint) error {
	err := us.qry.Deactive(uint(userID))
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// Profile implements users.UserService
// func (us *userService) Profile(userID uint) (users.Core, error) {
// 	res, err := us.qry.Profile(userID)
// 	if err != nil {
// 		log.Println("data not found")
// 		return users.Core{}, errors.New("query error, problem with server")
// 	}

// 	return res, nil
// }

// // Update implements users.UserService
// func (us *userService) Update(userID uint, fileData multipart.FileHeader, updateData users.Core) (users.Core, error) {
// 	url, err := helpers.GetUrlImagesFromAWS(fileData, int(1))
// 	if err != nil {
// 		return users.Core{}, errors.New("validate: " + err.Error())
// 	}
// 	updateData.ProfilePicture = url
// 	res, err := us.qry.Update(uint(userID), updateData)
// 	if err != nil {
// 		msg := ""
// 		if strings.Contains(err.Error(), "not found") {
// 			msg = "account not registered"
// 		} else if strings.Contains(err.Error(), "email") {
// 			msg = "email duplicated"
// 		} else if strings.Contains(err.Error(), "access denied") {
// 			msg = "access denied"
// 		}
// 		return users.Core{}, errors.New(msg)
// 	}
// 	return res, nil
// }
