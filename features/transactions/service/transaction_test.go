package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectTransactionsByUserId(t *testing.T) {
	table := SelectTransactionsByUserIdTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mTransaction := new(mocks.TransactionData_)
			mTransaction.On("SelectTransactionsByUserId", mock.Anything).Return(v.Output.results, v.Output.errResult)

			mPaymentGateway := new(mocks.PaymentGateway_)
			
			service := New(mTransaction, mPaymentGateway)
			_, err := service.SelectTransactionsByUserId(&v.Input.transactionEntity)
			if index <= 1 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestInsertTransaction(t *testing.T) {
	table := InsertTransactionTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mTransaction := new(mocks.TransactionData_)
			mTransaction.On("InsertTransaction", mock.Anything).Return(v.Output.result, v.Output.errResult)

			mPaymentGateway := new(mocks.PaymentGateway_)
			if index == 7 {
				mTransaction.On("ChargeTransaction", v.Output.result).Return(v.Output.result, v.Output.isErrChargeTransaction)
			} else {
				mTransaction.On("ChargeTransaction", v.Output.result).Return(v.Output.VirtualAccountNumber, v.Output.isErrChargeTransaction)
			}
			
			service := New(mTransaction, mPaymentGateway)
			_, err := service.InsertTransaction(&v.Input.transactionEntity)
			if index <= 4 {
				if v.Output.isErrValidate {
					assert.NotNil(t, err)
				} else if !v.Output.isErrValidate {
					assert.Nil(t, err)
				}
			} else if index >= 5 && index <= 6 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			} else if index >= 7 && index <= 8 {
				if v.Output.isErrChargeTransaction {
					assert.NotNil(t, err)
				} else if !v.Output.isErrChargeTransaction {
					assert.Nil(t, err)
				}
			} else if index >= 9 && index <= 10 {
				if v.Output.isErrorSecond {
					assert.NotNil(t, err)
				} else if !v.Output.isErrorSecond {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestUpdateTransactionById(t *testing.T) {
	table := UpdateTransactionStatusByIdTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mTransaction := new(mocks.TransactionData_)
			mTransaction.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mTransaction.On("UpdateTransactionStatusById", mock.Anything).Return(v.Output.errResult)
			mTransaction.On("SelectTransactionById", mock.Anything).Return(&v.Output.transactionGorm)

			mPaymentGateway := new(mocks.PaymentGateway_)
			
			service := New(mTransaction, mPaymentGateway)
			err := service.UpdateTransactionStatusById(&v.Input.transactionEntity)
			if index <= 1 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else if index >= 2 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

func TestUpdateTransactionByMidtrans(t *testing.T) {
	table := UpdateTransactionStatusByMidtransTestTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mTransaction := new(mocks.TransactionData_)
			mTransaction.On("UpdateTransactionStatusByMidtrans", mock.Anything).Return(v.Output.errResult)

			mPaymentGateway := new(mocks.PaymentGateway_)
			
			service := New(mTransaction, mPaymentGateway)
			err := service.UpdateTransactionStatusByMidtrans(&v.Input.transactionEntity)
			if index <= 1 {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}

// func TestUpdateCartById(t *testing.T) {
// 	table := UpdateCartByIdTestTable()
// 	for index, v := range table {
// 		t.Run(v.Name, func(t *testing.T) {
// 			mCart := new(mocks.CartData_)
// 			mCart.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
// 			mCart.On("UpdateCartById", mock.Anything).Return(v.Output.errResult)

// 			service := New(mCart)
// 			err := service.UpdateCartById(&v.Input.cartEntity)
// 			if index <= 1 {
// 				if v.Output.isErrValidate {
// 					assert.NotNil(t, err)
// 				} else if !v.Output.isErrValidate {
// 					assert.Nil(t, err)
// 				}
// 			} else if index >= 2 && index <= 3 {
// 				if !v.Output.isEntitled {
// 					assert.NotNil(t, err)
// 				} else if v.Output.isEntitled {
// 					assert.Nil(t, err)
// 				}
// 			} else if index >= 4 && index <= 5 {
// 				if v.Output.IsError {
// 					assert.NotNil(t, err)
// 				} else if !v.Output.IsError {
// 					assert.Nil(t, err)
// 				}
// 			}
// 		})
// 	}
// }