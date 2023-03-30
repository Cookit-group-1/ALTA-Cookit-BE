package service

import (
	"alta-cookit-be/features/transactions"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type TransactionService struct {
	transactionData transactions.TransactionData_
	validate        *validator.Validate
}

func New(transactionData transactions.TransactionData_) transactions.TransactionService_ {
	return &TransactionService{
		transactionData: transactionData,
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
	return output, nil
}

func (s *TransactionService) UpdateTransactionById(entity *transactions.TransactionEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.transactionData.ActionValidator(entity.ID, entity.CustomerUserId)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err = s.transactionData.UpdateTransactionById(entity)
	if err != nil {
		return err
	}
	return nil
}
