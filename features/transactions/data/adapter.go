package data

import (
	"alta-cookit-be/features/transaction_details"
	_transactionDetailData "alta-cookit-be/features/transaction_details/data"
	"alta-cookit-be/features/transactions"
	_transactionModel "alta-cookit-be/features/transactions/models"
)

func ConvertToGorm(entity *transactions.TransactionEntity) *_transactionModel.Transaction {
	gorm := _transactionModel.Transaction{
		UserID:            entity.CustomerUserId,
		TransactionDetail: *_transactionDetailData.ConvertToGorms(&entity.TransactionDetailEntities),
		IngredientID:      entity.IngredientID,
		Status:            entity.Status,
		PaymentMethod:     entity.PaymentMethod,
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

func ConvertToEntity(gorm *_transactionModel.Transaction, transactionDetailGorms *[]transaction_details.TransactionDetailEntity) *transactions.TransactionEntity {
	entity := transactions.TransactionEntity{
		ID:                        gorm.ID,
		TransactionDetailEntities: []transaction_details.TransactionDetailEntity{},
		CustomerUserId:            gorm.UserID,
		IngredientID:              gorm.IngredientID,
		Status:                    gorm.Status,
		PaymentMethod:             gorm.PaymentMethod,
		TotalPrice:                gorm.TotalPrice,
		CreatedAt:                 gorm.CreatedAt,
	}
	return &entity
}
