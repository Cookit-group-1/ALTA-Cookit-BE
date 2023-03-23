package delivery

import "alta-cookit-be/features/ingredient_details"

func ConvertToEntity(request *ingredient_details.IngredientDetailRequest) *ingredient_details.IngredientDetailEntity {
	return &ingredient_details.IngredientDetailEntity{
		ID:           request.ID,
		UserID:       request.UserID,
		UserRole:     request.UserRole,
		IngredientID: request.IngredientID,
		Name:         request.Name,
		Quantity:     request.Quantity,
		Unit:         request.Unit,
	}
}

func ConvertToResponse(entity *ingredient_details.IngredientDetailEntity) ingredient_details.IngredientDetailResponse {
	return ingredient_details.IngredientDetailResponse{
		ID:       entity.ID,
		Name:     entity.Name,
		Quantity: entity.Quantity,
		Unit:     entity.Unit,
	}
}
