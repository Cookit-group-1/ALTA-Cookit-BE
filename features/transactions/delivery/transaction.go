package delivery

import (
	"alta-cookit-be/features/transactions"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionDelivery struct {
	transactionService transactions.TransactionService_
}

func New(transactionService transactions.TransactionService_) transactions.TransactionDelivery_ {
	return &TransactionDelivery{
		transactionService: transactionService,
	}
}

func (d *TransactionDelivery) SelectTransactionByUserId(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	page, limit := helpers.ExtractPageLimit(e)
	limit, offset := helpers.LimitOffsetConvert(page, limit)

	transactionRequest := transactions.TransactionRequest{}
	err := e.Bind(&transactionRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	transactionRequest.CustomerUserId = userId
	transactionRequest.DataLimit = limit
	transactionRequest.DataOffset = offset

	outputs, err := d.transactionService.SelectTransactionByUserId(ConvertToEntity(&transactionRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.TRANSACTION_SuccessReadUserTransactionList, ConvertToResponses(outputs)))
}

func (d *TransactionDelivery) InsertTransaction(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)

	transactionRequest := transactions.TransactionRequest{}
	err := e.Bind(&transactionRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	transactionRequest.CustomerUserId = userId

	output, err := d.transactionService.InsertTransaction(ConvertToEntity(&transactionRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.TRANSACTION_SuccessInsertUserTransaction, ConvertToResponse(output)))
}

func (d *TransactionDelivery) UpdateTransactionById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_TransactionId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	transactionRequest := transactions.TransactionRequest{}
	err = e.Bind(&transactionRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	transactionRequest.ID = id
	transactionRequest.CustomerUserId = userId

	err = d.transactionService.UpdateTransactionById(ConvertToEntity(&transactionRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.Response(consts.CART_SuccessUpdateUserCart))
}
