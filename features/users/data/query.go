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
	res := User{}
	qry := uq.db.Delete(&res, userID)

	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}

	return nil
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

// UpgradeUser implements users.UserData
func (uq *UserQuery) UpgradeUser(userID uint, approvement users.Core) (users.Core, error) {
	input := CoreToModel(approvement)
	err := uq.db.Where("id = ?", userID).Updates(&input).Error
	if err != nil {
		log.Println("get user error : ", err.Error())
		return users.Core{}, errors.New("failed to upgrade to verifieduser")
	}

	return approvement, nil
}

// SearchUser implements users.UserData
func (uq *UserQuery) SearchUser(userID uint, quote string) ([]users.Core, error) {
	findRes := []User{}
	err := uq.db.Where("username LIKE ?", "%"+quote+"%").Find(&findRes).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []users.Core{}, errors.New("no user found")
	}
	res := []users.Core{}
	for i := 0; i < len(findRes); i++ {
		res = append(res, ModelToCore(findRes[i]))

	}
	if len(res) == 0 {
		return []users.Core{}, errors.New("no user found")
	}

	return res, nil
}

// ListUserRequest implements users.UserData
func (uq *UserQuery) ListUserRequest(userID uint) ([]users.Core, error) {
	listUser := []User{}
	err := uq.db.Where("approvement = ?", "requested").Find(&listUser).Error
	if err != nil {
		log.Println("get list user request query error", err.Error())
		return []users.Core{}, errors.New("users not found")
	}
	res := []users.Core{}
	for i := 0; i < len(listUser); i++ {
		res = append(res, ModelToCore(listUser[i]))

	}
	if len(res) == 0 {
		return []users.Core{}, errors.New("no user found")
	}
	return res, nil
}
