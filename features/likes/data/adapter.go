package data

import (
	"alta-cookit-be/features/likes"
	_likeModel "alta-cookit-be/features/likes/models"
)

func ConvertToGorm(entity *likes.LikeEntity) *_likeModel.Like {
	return &_likeModel.Like{
		UserID:   entity.UserID,
		RecipeID: entity.RecipeID,
	}
}
