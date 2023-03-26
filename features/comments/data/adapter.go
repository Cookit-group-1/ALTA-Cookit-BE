package data

import (
	"alta-cookit-be/features/comments"
	_commentModel "alta-cookit-be/features/comments/models"
	_userModel "alta-cookit-be/features/users/data"
)

func ConvertToGorm(entity *comments.CommentEntity) *_commentModel.Comment {
	gorm := _commentModel.Comment{
		UserID:   entity.UserID,
		RecipeID: entity.RecipeID,
		Comment:  entity.Comment,
		UrlImage: entity.UrlImage,
	}
	if entity.UrlImage == "" {
		gorm.UrlImage = " "
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToEntity(gorm *_commentModel.Comment, userGorm ...*_userModel.User) *comments.CommentEntity {
	entity := comments.CommentEntity{
		ID:       gorm.ID,
		UserID:   gorm.UserID,
		RecipeID: gorm.RecipeID,
		Comment:  gorm.Comment,
		UrlImage: gorm.UrlImage,
	}
	if len(userGorm) != 0 {
		entity.UserName = userGorm[0].Username
		entity.UserRole = userGorm[0].Role
		entity.ProfilePicture = userGorm[0].ProfilePicture
	}
	return &entity
}
