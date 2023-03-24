package delivery

import (
	_imageDelivery "alta-cookit-be/features/images/delivery"
	_ingredientsDelivery "alta-cookit-be/features/ingredients/delivery"
	"alta-cookit-be/features/recipes"
	_stepDelivery "alta-cookit-be/features/steps/delivery"
)

func ConvertToEntity(request *recipes.RecipeRequest) *recipes.RecipeEntity {
	return &recipes.RecipeEntity{
		ID:                 request.ID,
		UserID:             request.UserID,
		UserRole:           request.UserRole,
		RecipeID:           request.RecipeID,
		Image:              request.Image,
		ImageName:          request.ImageName,
		Type:               request.Type,
		Status:             request.Status,
		Name:               request.Name,
		Description:        request.Description,
		StepEntities:       *_stepDelivery.ConvertToEntities(&request.StepRequests),
		IngredientEntities: *_ingredientsDelivery.ConvertToEntities(&request.IngredientRequests),
	}
}

func ConvertToEntities(requests *[]recipes.RecipeRequest) *[]recipes.RecipeEntity {
	entities := []recipes.RecipeEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *recipes.RecipeEntity) *recipes.RecipeResponse {
	response := recipes.RecipeResponse{
		ID:                  entity.ID,
		UserID:              entity.UserID,
		UserName:            entity.UserName,
		UserRole:            entity.UserRole,
		ProfilePicture:      entity.ProfilePicture,
		RecipeID:            entity.RecipeID,
		Type:                entity.Type,
		Status:              entity.Status,
		Name:                entity.Name,
		Description:         entity.Description,
		StepResponses:       _stepDelivery.ConvertToResponses(&entity.StepEntities),
		IngredientResponses: _ingredientsDelivery.ConvertToResponses(&entity.IngredientEntities),
		ImageResponses:      _imageDelivery.ConvertToResponses(&entity.ImageEntities),
	}
	if entity.Recipe != nil {
		response.Recipe = ConvertToResponse(entity.Recipe)
	}
	return &response
}

func ConvertToResponses(entities *[]recipes.RecipeEntity) []recipes.RecipeResponse {
	responses := []recipes.RecipeResponse{}
	for _, entity := range *entities {
		responses = append(responses, *ConvertToResponse(&entity))
	}
	return responses
}
