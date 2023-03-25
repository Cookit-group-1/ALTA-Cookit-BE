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

func (uq *UserQuery) SelectUserById(user users.Core) *users.Core {
	existUser := CoreToModel(user)
	uq.db.Where("id = ?", existUser.ID).Find(&existUser)

	if existUser.Username == "" {
		return nil
	}

	user = ModelToCore(existUser)
	return &user
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

	newUser.ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
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
	if username == "" {
		log.Println("data empty, query error")
		return users.Core{}, errors.New("username is empty")
	}

	res := User{}
	if err := uq.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return users.Core{}, errors.New("data not found")
	}

	return ModelToCore(res), nil
}

// Deactive implements users.UserData
func (uq *UserQuery) Deactive(userID uint) error {
	panic("unimplemented")
}

// Profile implements users.UserData
func (uq *UserQuery) Profile(userID uint) (users.Core, error) {
	res := User{}
	err := uq.db.Where("id = ?", userID).First(&res).Error
	if err != nil {
		log.Println("get profile query error", err.Error())
		return users.Core{}, errors.New("account not found")
	}
	return ModelToCore(res), nil
}

// Update implements users.UserData
func (uq *UserQuery) Update(userID uint, updateData users.Core) (users.Core, error) {
	cnv := CoreToModel(updateData)
	res := User{}
	qry := uq.db.Model(&res).Where("id = ?", userID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return users.Core{}, errors.New("no data updated")
	}

	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return users.Core{}, err
	}

	return ModelToCore(cnv), nil
}
