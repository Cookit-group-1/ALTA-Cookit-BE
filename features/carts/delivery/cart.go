package delivery

import (
	"alta-cookit-be/features/carts"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartDeliery struct {
	cartService carts.CartService_
}

func New(cartService carts.CartService_) carts.CartDelivery_ {
	return &CartDeliery{
		cartService: cartService,
	}
}

func (d *CartDeliery) SelectCartsByUserId(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	page, limit := helpers.ExtractPageLimit(e)
	limit, offset := helpers.LimitOffsetConvert(page, limit)

	cartRequest := carts.CartRequest{}
	err := e.Bind(&cartRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	cartRequest.UserID = userId
	cartRequest.DataLimit = limit
	cartRequest.DataOffset = offset

	outputs, err := d.cartService.SelectCartsByUserId(ConvertToEntity(&cartRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.CART_SuccessInsertUserCart, ConvertToResponses(outputs)))
}

func (d *CartDeliery) InsertCart(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)

	cartRequest := carts.CartRequest{}
	err := e.Bind(&cartRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	cartRequest.UserID = userId

	output, err := d.cartService.InsertCart(ConvertToEntity(&cartRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.CART_SuccessInsertUserCart, ConvertToResponse(output)))
}

func (d *CartDeliery) UpdateCartById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_CartId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	cartRequest := carts.CartRequest{}
	err = e.Bind(&cartRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	cartRequest.ID = id
	cartRequest.UserID = userId

	err = d.cartService.UpdateCartById(ConvertToEntity(&cartRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.Response(consts.CART_SuccessUpdateUserCart))
}

func (d *CartDeliery) DeleteCartById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_CartId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	cartRequest := carts.CartRequest{}
	err = e.Bind(&cartRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	cartRequest.ID = id
	cartRequest.UserID = userId

	err = d.cartService.DeleteCartById(ConvertToEntity(&cartRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.Response(consts.CART_SuccessDeleteUserCart))
}
