package delivery

import (
	_imageDelivery "alta-cookit-be/features/images/delivery"
	"alta-cookit-be/features/transaction_details"
)

func ConvertToEntity(request *transaction_details.TransactionDetailRequest) *transaction_details.TransactionDetailEntity {
	return &transaction_details.TransactionDetailEntity{
		ID:             request.ID,
		LoggedInUserID: request.LoggedInUserID,
		IngredientID:   request.IngredientID,
		Quantity:       request.Quantity,
	}
}

func ConvertToEntities(requests *[]transaction_details.TransactionDetailRequest) *[]transaction_details.TransactionDetailEntity {
	entities := []transaction_details.TransactionDetailEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *transaction_details.TransactionDetailEntity) transaction_details.TransactionDetailResponse {
	return transaction_details.TransactionDetailResponse{
		ID:                   entity.ID,
		SellerUserID:         entity.SellerUserID,
		SellerUsername:       entity.SellerUsername,
		RecipeImageResponses: _imageDelivery.ConvertToResponses(&entity.RecipeImageEntities),
		RecipeName:           entity.RecipeName,
		IngredientName:       entity.IngredientName,
		Price:                entity.Price,
		Quantity:             entity.Quantity,
	}
}

func ConvertToResponses(entities *[]transaction_details.TransactionDetailEntity) []transaction_details.TransactionDetailResponse {
	responses := []transaction_details.TransactionDetailResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}
