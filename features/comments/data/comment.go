package data

import (
	"alta-cookit-be/app/storage"
	"alta-cookit-be/features/comments"
	_commentModel "alta-cookit-be/features/comments/models"
	"alta-cookit-be/utils/consts"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CommentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) comments.CommentData_ {
	return &CommentData{
		db: db,
	}
}

func (d *CommentData) SelectCommentById(id, recipe_id uint) *_commentModel.Comment {
	tempGorm := _commentModel.Comment{}
	d.db.Where("id = ? AND recipe_id = ?", id, recipe_id).Find(&tempGorm)
	
	if tempGorm.ID == 0 {
		return nil
	}
	return &tempGorm
}

func (d *CommentData) InsertComment (entity *comments.CommentEntity) (*comments.CommentEntity, error) {
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
	gorm, tempGorm := ConvertToGorm(entity), d.SelectCommentById(entity.ID, entity.RecipeID) 

	if tempGorm == nil {
		return nil, errors.New(consts.GORM_RecordNotFound)
	}
	if entity.Image != nil {
		err := storage.GetStorageClient().DeleteFile(tempGorm.UrlImage)
		if err != nil {
			return nil, err
		}

		urlImage, err := storage.GetStorageClient().UploadFile(entity.Image, entity.ImageName)
		if err != nil {
			return nil, err
		}
		gorm.UrlImage = urlImage
	}
	
	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Updates(ConvertToGorm(entity))
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

func (d *CommentData) DeleteCommentById(entity *comments.CommentEntity) error {
	tx := d.db.Where("id = ? AND recipe_id = ?", entity.ID, entity.RecipeID).Delete(ConvertToGorm(entity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0{
		return errors.New(consts.GORM_RecordNotFound)
	}
	return nil
}