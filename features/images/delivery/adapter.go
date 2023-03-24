package delivery

import (
	"alta-cookit-be/features/images"
)

func ConvertToEntity (request *images.ImageRequest) *images.ImageEntity {
	return &images.ImageEntity{
		ID:        request.ID,
		UserID:    request.UserID,
		RecipeID:  request.RecipeID,
		Image:     request.Image,
		ImageName: request.ImageName,
	}
}

func ConvertToEntities (requests *[]images.ImageRequest) *[]images.ImageEntity {
	entities := []images.ImageEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse (entity *images.ImageEntity) images.ImageResponse {
	return images.ImageResponse{
		ID:             entity.ID,
		UrlImage:       entity.UrlImage,
	}
}

func ConvertToResponses (entities *[]images.ImageEntity) []images.ImageResponse {
	responses := []images.ImageResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}