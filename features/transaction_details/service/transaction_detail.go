package service

import (
	"alta-cookit-be/features/transaction_details"

	"github.com/go-playground/validator"
)

type TransactionDetailService struct {
	transactionDetailData transaction_details.TransactionDetailData_
	validate              *validator.Validate
}

func New(transactionDetailData transaction_details.TransactionDetailData_) transaction_details.TransactionDetailService_ {
	return &TransactionDetailService{
		transactionDetailData: transactionDetailData,
		validate:              validator.New(),
	}
}

func (s *TransactionDetailService) SelectTransactionDetailById(entity *transaction_details.TransactionDetailEntity) (*transaction_details.TransactionDetailEntity, error) {
	outputs, err := s.transactionDetailData.SelectTransactionDetailById(entity)
	if err != nil {
		return nil, err
	}
	return outputs, nil
}
