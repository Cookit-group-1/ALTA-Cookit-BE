package service

import (
	"alta-cookit-be/features/transaction_details"
	"alta-cookit-be/features/transactions"
	_transactionModel "alta-cookit-be/features/transactions/models"
	"alta-cookit-be/utils/consts"
	"errors"

	"github.com/google/uuid"
)

type TestTable struct {
	Name  string
	Input struct {
		transactionEntity transactions.TransactionEntity
	}
	Output struct {
		isErrValidate          bool
		isEntitled             bool
		IsError                bool
		isErrChargeTransaction bool
		isErrorSecond          bool
		VirtualAccountNumber   string
		transactionGorm        _transactionModel.Transaction
		result                 *transactions.TransactionEntity
		results                *[]transactions.TransactionEntity
		errResult              error
	}
}

func SelectTransactionsByUserIdTable() []TestTable {
	tname := "test select transactions by user id "
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
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
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result:        nil,
				errResult:     nil,
			},
		},
	}
}

func InsertTransactionTestTable() []TestTable {
	tname := "test insert transaction "
	return []TestTable{
		{
			Name: tname + "expect failed - empty shipping fee",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty shipping method",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty transaction detail's ingredient id",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							Quantity: 1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty transaction detail's quantity",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect success - filled all required attribute",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					PaymentMethod: consts.TRANSACTION_E_BCA,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				isErrChargeTransaction: true,
				result:        &transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					PaymentMethod: consts.TRANSACTION_E_BCA,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
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
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					PaymentMethod: consts.TRANSACTION_E_BCA,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - failed charge transaction",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate:          false,
				isEntitled:             true,
				IsError:                false,
				isErrChargeTransaction: true,
				result: nil,
				errResult: errors.New(consts.SERVER_InternalServerError),
			},
		},
		{
			Name: tname + "expect success - success charge transaction",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate:          false,
				isEntitled:             true,
				IsError:                false,
				isErrChargeTransaction: false,
				VirtualAccountNumber:   "",
				result: &transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - second insert transaction",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate:          false,
				isEntitled:             true,
				IsError:                true,
				isErrChargeTransaction: false,
				VirtualAccountNumber:   "",
				isErrorSecond:          true,
				result: nil,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect failed - second insert transaction",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isErrValidate:          false,
				isEntitled:             true,
				IsError:                true,
				isErrChargeTransaction: false,
				VirtualAccountNumber:   "",
				isErrorSecond:          false,
				result: &transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					PaymentMethod:     consts.TRANSACTION_E_BCA,
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
							IngredientName: "Testing",
						},
					},
				},
				errResult: nil,
			},
		},
	}
}

func UpdateTransactionStatusByIdTestTable() []TestTable {
	tname := "test update transaction status by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    false,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed (status - unpaid)",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       true,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Unpaid,
					PaymentMethod: consts.TRANSACTION_E_BCA,
				},
				result: nil,
				errResult: errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success (status - unpaid)",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       false,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Unpaid,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: nil,
			},
		},
		{
			Name: tname + "expect success (status - shipped)",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       false,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Shipped,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: nil,
			},
		},
		{
			Name: tname + "expect success (status - received)",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       false,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Received,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       true,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Received,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: consts.TRANSACTION_E_Unpaid,
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       false,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Received,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: nil,
			},
		},
	}
}

func UpdateTransactionStatusByMidtransTestTable() []TestTable {
	tname := "test update transaction status by midtrans "
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: "settlement",
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       true,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Received,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				transactionEntity transactions.TransactionEntity
			}{
				transactionEntity: transactions.TransactionEntity{
					ID:                1,
					OrderID: (uuid.New()).String(),
					ShippingFee:       9000,
					ShippingMethod:    "GoFood",
					TransactionStatus: "settlement",
					TransactionDetailEntities: []transaction_details.TransactionDetailEntity{
						transaction_details.TransactionDetailEntity{
							IngredientID: 1,
							Quantity:     1,
						},
					},
				},
			},
			Output: struct {
				isErrValidate          bool
				isEntitled             bool
				IsError                bool
				isErrChargeTransaction bool
				isErrorSecond          bool
				VirtualAccountNumber   string
				transactionGorm        _transactionModel.Transaction
				result                 *transactions.TransactionEntity
				results                *[]transactions.TransactionEntity
				errResult              error
			}{
				isEntitled:    true,
				IsError:       false,
				transactionGorm: _transactionModel.Transaction{
					Status:        consts.TRANSACTION_E_Received,
					PaymentMethod: consts.TRANSACTION_E_COD,
				},
				result: nil,
				errResult: nil,
			},
		},
	}
}
