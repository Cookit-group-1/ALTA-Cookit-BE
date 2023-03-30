package delivery

import (
	"alta-cookit-be/features/carts"
)

func ConvertToEntity(request *carts.CartRequest) *carts.CartEntity {
	return &carts.CartEntity{
		ID:           request.ID,
		UserID:       request.UserID,
		IngredientID: request.IngredientID,
		Quantity:     request.Quantity,
	}
}

func ConvertToEntities(requests *[]carts.CartRequest) *[]carts.CartEntity {
	entities := []carts.CartEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *carts.CartEntity) carts.CartResponse {
	return carts.CartResponse{
		ID:             entity.ID,
		RecipeName:     entity.RecipeName,
		IngredientName: entity.IngredientName,
		Price:          entity.Price,
		Quantity:       entity.Quantity,
	}
}

func ConvertToResponses(entities *[]carts.CartEntity) []carts.CartResponse {
	responses := []carts.CartResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}
