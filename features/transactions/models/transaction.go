package models

import (
	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID             uint
	TransactionDetails []_transactionDetailModel.TransactionDetail `gorm:"constraint:OnDelete:CASCADE;"`
	Status             string                                      `gorm:"not null;type:enum('Unpaid', 'Shipped', 'Received', 'Complete');default:'Unpaid'"`
	PaymentMethod      string                                      `gorm:"not null;type:enum('None', 'COD', 'SeaBank', 'BCA', 'BNI', 'Mandiri', 'QRIS', 'Gopay');default:'None'"`
	TotalPrice         float64                                     `gorm:"not null"`
}
