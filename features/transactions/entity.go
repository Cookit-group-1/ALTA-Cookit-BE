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
	OrderID                   string
	TransactionStatus         string
	VirtualAccountNumber      string
	CustomerUserId            uint
	Status                    string
	PaymentMethod             string
	ShippingFee               float64 `validate:"required"`
	ShippingMethod            string  `validate:"required"`
	TotalPrice                float64
	CreatedAt                 time.Time
	DataLimit                 int
	DataOffset                int
	ExtractedQueryParams      map[string]interface{}
}

type TransactionRequest struct {
	ID                        uint                                           `json:"-" form:"-"`
	OrderID                   string                                         `json:"order_id,omitempty"`
	TransactionStatus         string                                         `json:"transaction_status" form:"transaction_status"`
	TransactionDetailRequests []transaction_details.TransactionDetailRequest `json:"transaction_details" form:"transaction_details"`
	CustomerUserId            uint                                           `json:"-" form:"-"`
	PaymentMethod             string                                         `json:"payment_method" form:"payment_method"`
	ShippingFee               float64                                        `json:"shipping_fee" form:"shipping_fee"`
	ShippingMethod            string                                         `json:"shipping_method" form:"shipping_method"`
	DataLimit                 int
	DataOffset                int
	ExtractedQueryParams      map[string]interface{}
}

type TransactionResponse struct {
	ID                         uint                                            `json:"id,omitempty"`
	OrderID                    string                                          `json:"order_id,omitempty"`
	VirtualAccountNumber       string                                          `json:"virtual_account_number,omitempty"`
	TransactionDetailResponses []transaction_details.TransactionDetailResponse `json:"transaction_details,omitempty"`
	CustomerUserId             uint                                            `json:"customer_id,omitempty"`
	Status                     string                                          `json:"status,omitempty"`
	PaymentMethod              string                                          `json:"payment_method,omitempty"`
	ShippingFee                float64                                         `json:"shipping_fee,omitempty"`
	ShippingMethod             string                                          `json:"shipping_method,omitempty"`
	TotalPrice                 float64                                         `json:"total_price,omitempty"`
	CreatedAt                  string                                          `json:"created_at,omitempty"`
}

type TransactionDelivery_ interface {
	SelectTransactionsByUserId(e echo.Context) error
	InsertTransaction(e echo.Context) error
	UpdateTransactionStatusById(e echo.Context) error
	UpdateTransactionStatusByOrderId(e echo.Context) error
}

type TransactionService_ interface {
	SelectTransactionsByUserId(transactionEntity *TransactionEntity) (*[]TransactionEntity, error)
	InsertTransaction(transactionEntity *TransactionEntity) (*TransactionEntity, error)
	UpdateTransactionStatusById(transactionEntity *TransactionEntity) error
	UpdateTransactionStatusByOrderId(transactionEntity *TransactionEntity) error
}

type TransactionData_ interface {
	ActionValidator(id, customerUserId uint) bool
	SelectTransactionById(id uint) *_transactionModel.Transaction
	SelectTransactionByTransactionDetailId(transactionDetailId uint) *_transactionModel.Transaction
	SelectTransactionsByUserId(transactionEntity *TransactionEntity) (*[]TransactionEntity, error)
	InsertTransaction(transactionEntity *TransactionEntity) (*TransactionEntity, error)
	UpdateTransactionStatusById(transactionEntity *TransactionEntity) error
	UpdateTransactionStatusByOrderId(transactionEntity *TransactionEntity) error
}
