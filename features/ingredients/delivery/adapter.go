package delivery

import "alta-cookit-be/features/ingredients"

func ConvertToEntity (request *ingredients.IngredientRequest) *ingredients.IngredientEntity {
	return &ingredients.IngredientEntity{
		ID:       request.ID,
		UserID:   request.UserID,
		UserRole: request.UserRole,
		RecipeID: request.RecipeID,
		Name:     request.Name,
		Price:    request.Price,
	}
}

func ConvertToResponse (entity *ingredients.IngredientEntity) ingredients.IngredientResponse {
	return ingredients.IngredientResponse{
		ID:    entity.ID,
		Name:  entity.Name,
		Price: entity.Price,
	}
}