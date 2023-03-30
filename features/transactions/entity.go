package transactions

import (
	"alta-cookit-be/features/transaction_details"
	"time"

	_transactionModel "alta-cookit-be/features/transactions/models"

	"github.com/labstack/echo/v4"
)

type TransactionEntity struct {
	ID                        uint
	TransactionDetailEntities []transaction_details.TransactionDetailEntity
	CustomerUserId            uint
	Status                    string
	PaymentMethod             string
	TotalPrice                float64
	CreatedAt                 time.Time
	DataLimit                 int
	DataOffset                int
}

type TransactionRequest struct {
	ID                        uint                                           `json:"-" form:"-"`
	TransactionDetailRequests []transaction_details.TransactionDetailRequest `json:"transaction_details" form:"transaction_details"`
	CustomerUserId            uint                                           `json:"-" form:"-"`
	Status                    string                                         `json:"status" form:"status"`
	PaymentMethod             string                                         `json:"payment_method" form:"payment_method"`
	DataLimit                 int
	DataOffset                int
}

type TransactionResponse struct {
	ID                         uint                                            `json:"id,omitempty"`
	TransactionDetailResponses []transaction_details.TransactionDetailResponse `json:"transaction_details,omitempty"`
	CustomerUserId             uint                                            `json:"customer_id,omitempty"`
	Status                     string                                          `json:"status,omitempty"`
	PaymentMethod              string                                          `json:"payment_method,omitempty"`
	TotalPrice                 float64                                         `json:"total_price,omitempty"`
	CreatedAt                  string                                          `json:"created_at,omitempty"`
}

type TransactionDelivery_ interface {
	SelectTransactionByUserId(e echo.Context) error
	InsertTransaction(e echo.Context) error
	UpdateTransactionById(e echo.Context) error
}

type TransactionService_ interface {
	SelectTransactionByUserId(transactionEntity *TransactionEntity) (*[]TransactionEntity, error)
	InsertTransaction(transactionEntity *TransactionEntity) (*TransactionEntity, error)
	UpdateTransactionById(transactionEntity *TransactionEntity) error
}

type TransactionData_ interface {
	ActionValidator(id, customerUserId uint) bool
	SelectTransactionByTransactionDetailId(transactionDetailId uint) *_transactionModel.Transaction
	SelectTransactionByUserId(transactionEntity *TransactionEntity) (*[]TransactionEntity, error)
	InsertTransaction(transactionEntity *TransactionEntity) (*TransactionEntity, error)
	UpdateTransactionById(transactionEntity *TransactionEntity) error
}
