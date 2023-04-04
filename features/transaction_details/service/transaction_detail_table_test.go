package service

import (
	"alta-cookit-be/features/transaction_details"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		transactionDetailEntity transaction_details.TransactionDetailEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *transaction_details.TransactionDetailEntity
		results       *[]transaction_details.TransactionDetailEntity
		errResult     error
	}
}

func SelectTransactionDetailByIdTable() []TestTable {
	tname := "test select transaction detail by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				transactionDetailEntity transaction_details.TransactionDetailEntity
			}{
				transactionDetailEntity: transaction_details.TransactionDetailEntity{
					ID:             1,
					LoggedInUserID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *transaction_details.TransactionDetailEntity
				results       *[]transaction_details.TransactionDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				transactionDetailEntity transaction_details.TransactionDetailEntity
			}{
				transactionDetailEntity: transaction_details.TransactionDetailEntity{
					ID:             1,
					LoggedInUserID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *transaction_details.TransactionDetailEntity
				results       *[]transaction_details.TransactionDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				transactionDetailEntity transaction_details.TransactionDetailEntity
			}{
				transactionDetailEntity: transaction_details.TransactionDetailEntity{
					ID:             1,
					LoggedInUserID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *transaction_details.TransactionDetailEntity
				results       *[]transaction_details.TransactionDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				transactionDetailEntity transaction_details.TransactionDetailEntity
			}{
				transactionDetailEntity: transaction_details.TransactionDetailEntity{
					ID:             1,
					LoggedInUserID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *transaction_details.TransactionDetailEntity
				results       *[]transaction_details.TransactionDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result:        &transaction_details.TransactionDetailEntity{
					ID:             1,
					LoggedInUserID: 1,
				},
				errResult:     nil,
			},
		},
	}
}
