package delivery

import (
	"alta-cookit-be/features/likes"
)

func ConvertToEntity(request *likes.LikeRequest) *likes.LikeEntity {
	return &likes.LikeEntity{
		UserID:   request.UserID,
		RecipeID: request.RecipeID,
	}
}
