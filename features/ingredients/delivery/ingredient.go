package delivery

import (
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IngredientDelivery struct {
	ingredientService ingredients.IngredientService_
}

func New(ingredientService ingredients.IngredientService_) ingredients.IngredientDelivery_ {
	return &IngredientDelivery{
		ingredientService: ingredientService,
	}
}

func (d *IngredientDelivery) InsertIngredient (e echo.Context) error {
	userId, userRole, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientRequest := ingredients.IngredientRequest{}
	err = e.Bind(&ingredientRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientRequest.RecipeID = recipeId
	ingredientRequest.UserID = userId
	ingredientRequest.UserRole = userRole

	output, err := d.ingredientService.InsertIngredient(ConvertToEntity(&ingredientRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.INGREDIENT_SuccessInsertRecipeIngredient, ConvertToResponse(output)))
}

func (d *IngredientDelivery) UpdateIngredientById(e echo.Context) error {
	userId, userRole, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientRequest := ingredients.IngredientRequest{}
	err = e.Bind(&ingredientRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientRequest.ID = id
	ingredientRequest.UserID = userId
	ingredientRequest.RecipeID = recipeId
	ingredientRequest.UserRole = userRole

	err = d.ingredientService.UpdateIngredientById(ConvertToEntity(&ingredientRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.INGREDIENT_SuccessUpdateRecipeIngredient))
}

func (d *IngredientDelivery) DeleteIngredientById(e echo.Context) error {
	userId, userRole, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_IngredientId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientRequest := ingredients.IngredientRequest{}
	err = e.Bind(&ingredientRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientRequest.ID = id
	ingredientRequest.UserID = userId
	ingredientRequest.RecipeID = recipeId
	ingredientRequest.UserRole = userRole

	err = d.ingredientService.DeleteIngredientById(ConvertToEntity(&ingredientRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.INGREDIENT_SuccessDeleteRecipeIngredient))
}

func (d *IngredientDelivery) DeleteIngredientByRecipeId(e echo.Context) error {
	userId, userRole, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	ingredientRequest := ingredients.IngredientRequest{}
	err = e.Bind(&ingredientRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	ingredientRequest.UserID = userId
	ingredientRequest.RecipeID = recipeId
	ingredientRequest.UserRole = userRole

	err = d.ingredientService.DeleteIngredientByRecipeId(ConvertToEntity(&ingredientRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.INGREDIENT_SuccessDeleteRecipeIngredient))
}