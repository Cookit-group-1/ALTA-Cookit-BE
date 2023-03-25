package delivery

import (
	"alta-cookit-be/features/steps"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StepDelivery struct {
	stepService steps.StepService_
}

func New(stepService steps.StepService_) steps.StepDelivery_ {
	return &StepDelivery{
		stepService: stepService,
	}
}

func (d *StepDelivery) InsertStep(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	stepRequest := steps.StepRequest{}
	err = e.Bind(&stepRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	stepRequest.RecipeID = recipeId
	stepRequest.UserID = userId

	output, err := d.stepService.InsertStep(ConvertToEntity(&stepRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.STEP_SuccesInsertRecipeStep, ConvertToResponse(output)))
}

func (d *StepDelivery) UpdateStepById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_StepId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	stepRequest := steps.StepRequest{}
	err = e.Bind(&stepRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	stepRequest.ID = id
	stepRequest.RecipeID = recipeId
	stepRequest.UserID = userId

	err = d.stepService.UpdateStepById(ConvertToEntity(&stepRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.STEP_SuccesUpdateRecipeStep))
}

func (d *StepDelivery) DeleteStepById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_StepId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	stepRequest := steps.StepRequest{}
	err = e.Bind(&stepRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	stepRequest.ID = id
	stepRequest.UserID = userId
	stepRequest.RecipeID = recipeId

	err = d.stepService.DeleteStepById(ConvertToEntity(&stepRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.STEP_SuccessDeleteRecipeStep))
}
