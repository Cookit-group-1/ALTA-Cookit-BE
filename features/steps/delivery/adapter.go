package delivery

import (
	"alta-cookit-be/features/steps"
)

func ConvertToEntity(request *steps.StepRequest) *steps.StepEntity {
	return &steps.StepEntity{
		ID:       request.ID,
		UserID:   request.UserID,
		UserRole: request.UserRole,
		RecipeID: request.RecipeID,
		Name:     request.Name,
	}
}

func ConvertToEntities(requests *[]steps.StepRequest) *[]steps.StepEntity {
	entities := []steps.StepEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *steps.StepEntity) steps.StepResponse {
	return steps.StepResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func ConvertToResponses(entities *[]steps.StepEntity) []steps.StepResponse {
	responses := []steps.StepResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}
