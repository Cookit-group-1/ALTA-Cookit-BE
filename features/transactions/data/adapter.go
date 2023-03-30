package data

import (
	"alta-cookit-be/features/transaction_details"
	_transactionDetailData "alta-cookit-be/features/transaction_details/data"
	"alta-cookit-be/features/transactions"
	_transactionModel "alta-cookit-be/features/transactions/models"
)

func ConvertToGorm(entity *transactions.TransactionEntity) *_transactionModel.Transaction {
	gorm := _transactionModel.Transaction{
		UserID:             entity.CustomerUserId,
		TransactionDetails: *_transactionDetailData.ConvertToGorms(&entity.TransactionDetailEntities),
		Status:             entity.Status,
		PaymentMethod:      entity.PaymentMethod,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]transactions.TransactionEntity) *[]_transactionModel.Transaction {
	gorms := []_transactionModel.Transaction{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_transactionModel.Transaction, transactionDetailEntities *[]transaction_details.TransactionDetailEntity) *transactions.TransactionEntity {
	entity := transactions.TransactionEntity{
		ID:                        gorm.ID,
		TransactionDetailEntities: *transactionDetailEntities,
		CustomerUserId:            gorm.UserID,
		Status:                    gorm.Status,
		PaymentMethod:             gorm.PaymentMethod,
		TotalPrice:                gorm.TotalPrice,
		CreatedAt:                 gorm.CreatedAt,
	}
	return &entity
}
