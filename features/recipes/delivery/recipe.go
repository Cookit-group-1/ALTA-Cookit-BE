package delivery

import (
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
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

	output, err := d.recipeService.InsertRecipe(ConvertToEntity(&recipeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RECIPE_SuccessInsertUserRecipe, ConvertToResponse(output)))
}
