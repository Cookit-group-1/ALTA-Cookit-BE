package delivery

import "alta-cookit-be/features/ingredient_details"

func ConvertToEntity(request *ingredient_details.IngredientDetailRequest) *ingredient_details.IngredientDetailEntity {
	return &ingredient_details.IngredientDetailEntity{
		ID:           request.ID,
		UserID:       request.UserID,
		RecipeID:     request.RecipeID,
		IngredientID: request.IngredientID,
		Name:         request.Name,
		Quantity:     request.Quantity,
		Unit:         request.Unit,
	}
}

func ConvertToEntities(requests *[]ingredient_details.IngredientDetailRequest) *[]ingredient_details.IngredientDetailEntity {
	entities := []ingredient_details.IngredientDetailEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *ingredient_details.IngredientDetailEntity) ingredient_details.IngredientDetailResponse {
	return ingredient_details.IngredientDetailResponse{
		ID:       entity.ID,
		Name:     entity.Name,
		Quantity: entity.Quantity,
		Unit:     entity.Unit,
	}
}

func ConvertToResponses(entities *[]ingredient_details.IngredientDetailEntity) []ingredient_details.IngredientDetailResponse {
	responses := []ingredient_details.IngredientDetailResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}
