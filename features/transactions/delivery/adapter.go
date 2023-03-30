package delivery

import (
	_transactionDetailDelivery "alta-cookit-be/features/transaction_details/delivery"
	"alta-cookit-be/features/transactions"
)

func ConvertToEntity(request *transactions.TransactionRequest) *transactions.TransactionEntity {
	return &transactions.TransactionEntity{
		ID:                        request.ID,
		TransactionDetailEntities: *_transactionDetailDelivery.ConvertToEntities(&request.TransactionDetailRequests),
		CustomerUserId:            request.CustomerUserId,
		Status:                    request.Status,
		PaymentMethod:             request.PaymentMethod,
		DataLimit:                 request.DataLimit,
		DataOffset:                request.DataOffset,
		ExtractedQueryParams:      request.ExtractedQueryParams,
	}
}

func ConvertToEntities(requests *[]transactions.TransactionRequest) *[]transactions.TransactionEntity {
	entities := []transactions.TransactionEntity{}
	for _, request := range *requests {
		entities = append(entities, *ConvertToEntity(&request))
	}
	return &entities
}

func ConvertToResponse(entity *transactions.TransactionEntity) transactions.TransactionResponse {
	return transactions.TransactionResponse{
		ID:                         entity.ID,
		TransactionDetailResponses: _transactionDetailDelivery.ConvertToResponses(&entity.TransactionDetailEntities),
		CustomerUserId:             entity.CustomerUserId,
		Status:                     entity.Status,
		PaymentMethod:              entity.PaymentMethod,
		TotalPrice:                 entity.TotalPrice,
		CreatedAt:                  entity.CreatedAt.Format("2006-01-02"),
	}
}

func ConvertToResponses(entities *[]transactions.TransactionEntity) []transactions.TransactionResponse {
	responses := []transactions.TransactionResponse{}
	for _, entity := range *entities {
		responses = append(responses, ConvertToResponse(&entity))
	}
	return responses
}
