package data

import (
	"alta-cookit-be/features/users"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &UserQuery{
		db: db,
	}
}

// Register implements users.UserData
func (uq *UserQuery) Register(newUser users.Core) (users.Core, error) {
	if newUser.Username == "" || newUser.Password == "" {
		log.Println("data empty")
		return users.Core{}, errors.New("username or password is empty")
	}

	dupUser := CoreToModel(newUser)
	err := uq.db.Where("username = ?", newUser.Username).First(&dupUser).Error
	if err == nil {
		log.Println("duplicated")
		return users.Core{}, errors.New("username duplicated")
	}

	cnv := CoreToModel(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

// Login implements users.UserData
func (uq *UserQuery) Login(username string) (users.Core, error) {
	panic("unimplemented")
}
