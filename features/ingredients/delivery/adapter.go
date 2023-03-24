package delivery

import (
	"alta-cookit-be/features/ingredients"
	_ingredientDetailDelivery "alta-cookit-be/features/ingredient_details/delivery"
)

func ConvertToEntity (request *ingredients.IngredientRequest) *ingredients.IngredientEntity {
	return &ingredients.IngredientEntity{
		ID:       request.ID,
		UserID:   request.UserID,
		UserRole: request.UserRole,
		RecipeID: request.RecipeID,
		Name:     request.Name,
		Price:    request.Price,
		IngredientDetailEntities: *_ingredientDetailDelivery.ConvertToEntities(&request.IngredientDetailRequests),
	}
}

func ConvertToEntities (requests *[]ingredients.IngredientRequest) *[]ingredients.IngredientEntity {
	entities := []ingredients.IngredientEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse (entity *ingredients.IngredientEntity) ingredients.IngredientResponse {
	return ingredients.IngredientResponse{
		ID:    entity.ID,
		Name:  entity.Name,
		Price: entity.Price,
		IngredientDetailResponses: _ingredientDetailDelivery.ConvertToResponses(&entity.IngredientDetailEntities),
	}
}

func ConvertToResponses (entities *[]ingredients.IngredientEntity) []ingredients.IngredientResponse {
	responses := []ingredients.IngredientResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}