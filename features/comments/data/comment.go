package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/comments"
	_commentModel "alta-cookit-be/features/comments/models"
	"alta-cookit-be/features/users"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CommentData struct {
	db *gorm.DB
	userData users.UserData_
}

func New(db *gorm.DB, userData users.UserData_) comments.CommentData_ {
	return &CommentData{
		db: db,
		userData: userData,
	}
}

func (d *CommentData) SelectCommentById(id uint) *_commentModel.Comment {
	tempGorm := _commentModel.Comment{}
	d.db.Where("id = ?", id).Find(&tempGorm)
	
	if tempGorm.ID == 0 {
		return nil
	}
	return &tempGorm
}

func (d *CommentData) SelectCommentsByRecipeId(entity *comments.CommentEntity) (*[]comments.CommentEntity, error) {
	gorms := []_commentModel.Comment{}
	
	tx := d.db.Where("recipe_id = ?", entity.RecipeID).Find(&gorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	entities := []comments.CommentEntity{}
	for _, comment := range gorms {
		userGorm := d.userData.SelectUserById(comment.UserID)
		entities = append(entities, *ConvertToEntity(&comment, userGorm))
	}

	return &entities, nil
}

func (d *CommentData) InsertComment(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	gorm := ConvertToGorm(entity)

	if entity.Image != nil {
		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)
		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}
	
	tx := d.db.Create(gorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		return nil, tx.Error
	}
	return ConvertToEntity(gorm), nil
}

func (d *CommentData) UpdateCommentById(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	gorm, tempGorm := ConvertToGorm(entity), d.SelectCommentById(entity.ID) 

	if tempGorm == nil {
		return nil, errors.New(consts.GORM_RecordNotFound)
	}
	if entity.Image != nil {
		if tempGorm.UrlImage != "" {
			err := storage.GetStorageClient().DeleteFile(tempGorm.UrlImage)
			if err != nil {
				return nil, err
			}
		}

		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)
		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}
	
	tx := d.db.Where("id = ?", entity.ID).Updates(ConvertToGorm(entity))
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "user_id") {
			return nil, errors.New(consts.USER_InvalidUser)
		}
		if strings.Contains(tx.Error.Error(), "recipe_id") {
			return nil, errors.New(consts.RECIPE_InvalidRecipe)
		}
		return nil, tx.Error
	}
	if tx.RowsAffected == 0{
		return nil, errors.New(consts.GORM_RecordNotFound)
	}
	return ConvertToEntity(gorm), nil
}

func (d *CommentData) DeleteCommentById(entity *comments.CommentEntity) error {
	tx := d.db.Where("id = ?", entity.ID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}