package delivery

import (
	"alta-cookit-be/features/ingredient_details"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IngredientDetailDelivery struct {
	ingredientDetailService ingredient_details.IngredientDetailService_
}

func New(ingredientDetailService ingredient_details.IngredientDetailService_) ingredient_details.IngredientDetailDelivery_ {
	return &IngredientDetailDelivery{
		ingredientDetailService: ingredientDetailService,
	}
}

func (d *IngredientDetailDelivery) InsertIngredientDetail (e echo.Context) error {
	ingredientId, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientDetailRequest := ingredient_details.IngredientDetailRequest{}
	err = e.Bind(&ingredientDetailRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientDetailRequest.IngredientID = ingredientId

	output, err := d.ingredientDetailService.InsertIngredientDetail(ConvertToEntity(&ingredientDetailRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.INGREDIENT_DETAIL_SuccessInsertIngredientDetail, ConvertToResponse(output)))
}

func (d *IngredientDetailDelivery) UpdateIngredientDetailById (e echo.Context) error {
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientDetailId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientId, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientDetailRequest := ingredient_details.IngredientDetailRequest{}
	err = e.Bind(&ingredientDetailRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientDetailRequest.ID = id
	ingredientDetailRequest.IngredientID = ingredientId

	err = d.ingredientDetailService.UpdateIngredientDetailById(ConvertToEntity(&ingredientDetailRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.INGREDIENT_DETAIL_SuccessUpdateIngredientDetail))
}

func (d *IngredientDetailDelivery) DeleteIngredientDetailById(e echo.Context) error {
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientDetailId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientId, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientDetailRequest := ingredient_details.IngredientDetailRequest{}
	err = e.Bind(&ingredientDetailRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientDetailRequest.ID = id
	ingredientDetailRequest.IngredientID = ingredientId

	err = d.ingredientDetailService.DeleteIngredientDetailById(ConvertToEntity(&ingredientDetailRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.INGREDIENT_DETAIL_SuccessDeleteIngredientDetail))
}