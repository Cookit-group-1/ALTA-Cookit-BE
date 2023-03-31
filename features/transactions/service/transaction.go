package service

import (
	"alta-cookit-be/app/payment"
	"alta-cookit-be/features/transactions"
	"alta-cookit-be/utils/consts"
	"strings"

	"errors"

	"github.com/go-playground/validator"
)

type TransactionService struct {
	transactionData transactions.TransactionData_
	paymentGateway  payment.PaymentGateway
	validate        *validator.Validate
}

func New(transactionData transactions.TransactionData_) transactions.TransactionService_ {
	return &TransactionService{
		transactionData: transactionData,
		paymentGateway:  payment.NewCoreMidtrans(),
		validate:        validator.New(),
	}
}

func (s *TransactionService) SelectTransactionsByUserId(entity *transactions.TransactionEntity) (*[]transactions.TransactionEntity, error) {
	outputs, err := s.transactionData.SelectTransactionsByUserId(entity)
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (s *TransactionService) InsertTransaction(entity *transactions.TransactionEntity) (*transactions.TransactionEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	output, err := s.transactionData.InsertTransaction(entity)
	if err != nil {
		return nil, err
	}

	if output.PaymentMethod != consts.TRANSACTION_E_NONE && output.PaymentMethod != consts.TRANSACTION_E_COD {
		viruatlAccountNumber, err := s.paymentGateway.ChargeTransaction(output.OrderID, strings.ToLower(output.PaymentMethod), output.TransactionDetailEntities)
		if err != nil {
			return nil, errors.New(consts.SERVER_InternalServerError)
		}

		output.VirtualAccountNumber = viruatlAccountNumber
	}
	return output, nil
}

func (s *TransactionService) UpdateTransactionStatusById(entity *transactions.TransactionEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.transactionData.ActionValidator(entity.ID, entity.CustomerUserId)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	gorm := s.transactionData.SelectTransactionById(entity.ID)
	switch gorm.Status {
		case consts.TRANSACTION_E_Unpaid:
			return errors.New(consts.SERVER_ForbiddenRequest)
		case consts.TRANSACTION_E_Shipped:
			entity.Status = consts.TRANSACTION_E_Received
		case consts.TRANSACTION_E_Received:
			entity.Status = consts.TRANSACTION_E_Complete
		default:
			return nil
	}
			
	err = s.transactionData.UpdateTransactionStatusById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) UpdateTransactionStatusByOrderId(entity *transactions.TransactionEntity) error {
	if entity.TransactionStatus == "settlement" {
		entity.Status = consts.TRANSACTION_E_Shipped
		err := s.transactionData.UpdateTransactionStatusByOrderId(entity)
		if err != nil {
			return err
		}
	}
	return nil
}
