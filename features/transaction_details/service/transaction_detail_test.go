package service

import (
	"alta-cookit-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectTransactionDetailById(t *testing.T) {
	table := SelectTransactionDetailByIdTable()
	for index, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			mTransactionDetail := new(mocks.TransactionDetailData_)
			mTransactionDetail.On("ActionValidator", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(v.Output.isEntitled)
			mTransactionDetail.On("SelectTransactionDetailById", mock.Anything).Return(v.Output.result, v.Output.errResult)

			service := New(mTransactionDetail)
			_, err := service.SelectTransactionDetailById(&v.Input.transactionDetailEntity)
			if index <= 1 {
				if !v.Output.isEntitled {
					assert.NotNil(t, err)
				} else if v.Output.isEntitled {
					assert.Nil(t, err)
				}
			} else {
				if v.Output.IsError {
					assert.NotNil(t, err)
				} else if !v.Output.IsError {
					assert.Nil(t, err)
				}
			}
		})
	}
}
