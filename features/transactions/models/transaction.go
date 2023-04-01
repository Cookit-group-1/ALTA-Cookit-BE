package models

import (
	_transactionDetailModel "alta-cookit-be/features/transaction_details/models"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID            string `gorm:"default:'';not null;type:VARCHAR(50)"`
	UserID             uint
	TransactionDetails []_transactionDetailModel.TransactionDetail `gorm:"constraint:OnDelete:CASCADE;"`
	Status             string                                      `gorm:"not null;type:enum('Unpaid', 'Shipped', 'Received', 'Complete');default:'Unpaid'"`
	PaymentMethod      string                                      `gorm:"not null;type:enum('None', 'COD', 'BCA', 'BNI', 'BRI', 'Mandiri', 'Permata');default:'None'"`
	ShippingFee        float64                                     `gorm:"not null"`
	ShippingMethod     string                                      `gorm:"default:''"`
	TotalPrice         float64                                     `gorm:"not null"`
}
