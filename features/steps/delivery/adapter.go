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

func ConvertToResponse(entity *steps.StepEntity) steps.StepResponse {
	return steps.StepResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
