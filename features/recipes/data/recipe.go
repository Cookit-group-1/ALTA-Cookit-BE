package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/images"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/users"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	_imageModel "alta-cookit-be/features/images/models"

	"gorm.io/gorm"
)

type RecipeData struct {
	db        *gorm.DB
	userData  users.UserData_
	imageData images.ImageData_
}

func New(db *gorm.DB, userData users.UserData_, imageData images.ImageData_) recipes.RecipeData_ {
	return &RecipeData{
		db:        db,
		userData:  userData,
		imageData: imageData,
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

func (d *RecipeData) UpdateRecipeById(entity *recipes.RecipeEntity) error {
	gorm := ConvertToGorm(entity)

	tx := d.db.Omit("recipe_id").Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Updates(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}

func (d *RecipeData) DeleteRecipeById(entity *recipes.RecipeEntity) error {
	gorm, imageGorms := ConvertToGorm(entity), d.imageData.SelectImagesByRecipeId(entity.ID)

	for _, imageGorm := range *imageGorms {
		err := storage.GetStorageClient().DeleteFile(imageGorm.UrlImage)
		if err != nil {
			return err
		}

		tx := d.db.Delete(&imageGorm)
		if tx.Error != nil {
			return tx.Error
		}
	}

	tx := d.db.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Delete(&gorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}
