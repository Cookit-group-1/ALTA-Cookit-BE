package data

import (
	"alta-cookit-be/features/comments"
	_commentModel "alta-cookit-be/features/comments/models"
)

func ConvertToGorm(entity *comments.CommentEntity) *_commentModel.Comment {
	gorm := _commentModel.Comment{
		UserID:   entity.UserID,
		RecipeID: entity.RecipeID,
		Comment:  entity.Comment,
		UrlImage: entity.UrlImage,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToEntity(gorm *_commentModel.Comment) *comments.CommentEntity {
	return &comments.CommentEntity{
		ID:       gorm.ID,
		UserID:   gorm.UserID,
		RecipeID: gorm.RecipeID,
		Comment:  gorm.Comment,
		UrlImage: gorm.UrlImage,
	}
}
