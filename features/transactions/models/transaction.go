package models

import (
	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID            uint
	TransactionDetail []_transactionDetailModel.TransactionsDetail
	IngredientID      uint
	Status            string  `gorm:"not null;type:enum('Unpaid', 'Shipped', 'Received', Complete);default:'Unpaid'"`
	PaymentMethod     string  `gorm:"not null;type:enum('COD', 'SeaBank', 'BCA', 'BNI', 'Mandiri', 'QRIS', 'Gopay');default:'Unpaid'"`
	TotalPrice        float64 `gorm:"not null"`
}
