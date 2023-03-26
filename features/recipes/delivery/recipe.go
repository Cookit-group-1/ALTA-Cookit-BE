package delivery

import (
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeDelivery struct {
	recipeService recipes.RecipeService_
}

func New(recipeService recipes.RecipeService_) recipes.RecipeDelivery_ {
	return &RecipeDelivery{
		recipeService: recipeService,
	}
}

func (d *RecipeDelivery) InsertRecipe(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeRequest := recipes.RecipeRequest{}
	err := e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}

	files, fileNames, _ := helpers.ExtractMultipleFiles(e, "image")
	for index, file := range files {
		recipeRequest.Image = append(recipeRequest.Image, file)
		recipeRequest.ImageName = append(recipeRequest.ImageName, fileNames[index])
	}
	recipeRequest.UserID = userId

	output, err := d.recipeService.InsertRecipe(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RECIPE_SuccessInsertUserRecipe, ConvertToResponse(output)))
}

func (d *RecipeDelivery) UpdateRecipeById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	recipeRequest := recipes.RecipeRequest{}
	err = e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	recipeRequest.ID = recipeId
	recipeRequest.UserID = userId

	err = d.recipeService.UpdateRecipeById(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.Response(consts.RECIPE_SuccessUpdateUserRecipe))
}

func (d *RecipeDelivery) DeleteRecipeById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	recipeRequest := recipes.RecipeRequest{}
	err = e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	recipeRequest.ID = recipeId
	recipeRequest.UserID = userId

	err = d.recipeService.DeleteRecipeById(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.Response(consts.RECIPE_SucessDeleteUserRecipe))
}

func (d *RecipeDelivery) SelectRecipesByUserId(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	page, limit, err := helpers.ExtractPageLimit(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}
	limit, offset := helpers.LimitOffsetConvert(page, limit)

	recipeRequest := recipes.RecipeRequest{}
	err = e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	recipeRequest.UserID = userId
	recipeRequest.DataLimit = limit
	recipeRequest.DataOffset = offset
	recipeRequest.ExtractedQueryParams = helpers.ExtractQueryParams(e.QueryParams())

	outputs, err := d.recipeService.SelectRecipesByUserId(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RECIPE_SuccessReadListOfUserRecipes, ConvertToResponses(outputs)))
}

func (d *RecipeDelivery) SelectRecipesTimeline(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	page, limit, err := helpers.ExtractPageLimit(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}
	limit, offset := helpers.LimitOffsetConvert(page, limit)

	recipeRequest := recipes.RecipeRequest{}
	err = e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	recipeRequest.UserID = userId
	recipeRequest.DataLimit = limit
	recipeRequest.DataOffset = offset
	recipeRequest.ExtractedQueryParams = helpers.ExtractQueryParams(e.QueryParams())

	outputs, err := d.recipeService.SelectRecipesTimeline(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RECIPE_SuccessReadListOfRecipes, ConvertToResponses(outputs)))
}

func (d *RecipeDelivery) SelectRecipeDetailById(e echo.Context) error {
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	recipeRequest := recipes.RecipeRequest{}
	err = e.Bind(&recipeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	recipeRequest.ID = recipeId

	output, err := d.recipeService.SelectRecipeDetailById(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RECIPE_SuccessReadDetailsOfRecipe, ConvertToResponse(output)))
}
