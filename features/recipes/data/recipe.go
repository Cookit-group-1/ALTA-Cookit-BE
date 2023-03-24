package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/users"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	_imageModel "alta-cookit-be/features/images/models"

	"gorm.io/gorm"
)

type RecipeData struct {
	db       *gorm.DB
	userData users.UserData_
}

func New(db *gorm.DB, userData users.UserData_) recipes.RecipeData_ {
	return &RecipeData{
		db:       db,
		userData: userData,
	}
}

func (d *RecipeData) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	gorm := *ConvertToGorm(entity)

	txTransaction := d.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	tx := txTransaction.Omit("recipe_id").Create(&gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		return nil, tx.Error
	}

	for index, file := range entity.Image {
		urlImage, err := storage.GetStorageClient().UploadFile(file, entity.ImageName[index])
		if err != nil {
			return nil, err
		}
		gorm.Images = append(gorm.Images, _imageModel.Image{
			UrlImage: urlImage,
		})
	}

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	userGorm := d.userData.SelectUserById(entity.UserID)
	return ConvertToEntity(&gorm, userGorm), nil
}
