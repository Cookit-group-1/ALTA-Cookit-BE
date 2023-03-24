package delivery

import (
	"alta-cookit-be/features/comments"
)

func ConvertToEntity (request *comments.CommentRequest) *comments.CommentEntity {
	return &comments.CommentEntity{
		ID:        request.ID,
		UserID:    request.UserID,
		UserRole:  request.UserRole,
		RecipeID:  request.RecipeID,
		Comment:   request.Comment,
		Image:     request.Image,
		ImageName: request.ImageName,
	}
}

func ConvertToResponse (entity *comments.CommentEntity) comments.CommentResponse {
	return comments.CommentResponse{
		ID:             entity.ID,
		UserName:       entity.UserName,
		UserRole:       entity.UserRole,
		ProfilePicture: entity.ProfilePicture,
		Comment:        entity.Comment,
		UrlImage:       entity.UrlImage,
	}
}

func ConvertToResponses (entities *[]comments.CommentEntity) []comments.CommentResponse {
	responses := []comments.CommentResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}