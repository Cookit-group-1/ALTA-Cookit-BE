package delivery

import (
	"alta-cookit-be/features/transaction_details"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionDetailDelivery struct {
	transactionDetailService transaction_details.TransactionDetailService_
}

func New(transactionDetailService transaction_details.TransactionDetailService_) transaction_details.TransactionDetailDelivery_ {
	return &TransactionDetailDelivery{
		transactionDetailService: transactionDetailService,
	}
}

func (d *TransactionDetailDelivery) SelectTransactionDetailById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_TransactionDetailId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	transactionDetailRequest := transaction_details.TransactionDetailRequest{}
	err = e.Bind(&transactionDetailRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	transactionDetailRequest.ID = id
	transactionDetailRequest.LoggedInUserID = userId

	output, err := d.transactionDetailService.SelectTransactionDetailById(ConvertToEntity(&transactionDetailRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.ResponseWithData(consts.TRANSACTION_DETAIL_SuccessReadUserTransactionDetailList, ConvertToResponse(output)))
}
