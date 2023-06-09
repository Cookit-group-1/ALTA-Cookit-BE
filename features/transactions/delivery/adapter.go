package delivery

import (
	_transactionDetailDelivery "alta-cookit-be/features/transaction_details/delivery"
	"alta-cookit-be/features/transactions"
)

func ConvertToEntity(request *transactions.TransactionRequest) *transactions.TransactionEntity {
	return &transactions.TransactionEntity{
		ID:                        request.ID,
		TransactionDetailEntities: *_transactionDetailDelivery.ConvertToEntities(&request.TransactionDetailRequests),
		OrderID:                   request.OrderID,
		TransactionStatus:         request.TransactionStatus,
		CustomerUserId:            request.CustomerUserId,
		PaymentMethod:             request.PaymentMethod,
		ShippingFee:               request.ShippingFee,
		ShippingMethod:            request.ShippingMethod,
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
		OrderID:                    entity.OrderID,
		VirtualAccountNumber:       entity.VirtualAccountNumber,
		CustomerUserId:             entity.CustomerUserId,
		Status:                     entity.Status,
		PaymentMethod:              entity.PaymentMethod,
		ShippingFee:                entity.ShippingFee,
		ShippingMethod:             entity.ShippingMethod,
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
