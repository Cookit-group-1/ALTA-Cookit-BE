package delivery

import (
	"alta-cookit-be/features/likes"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeDelivery struct {
	likeService likes.LikeService_
}

func New(likeService likes.LikeService_) likes.LikeDelivery_ {
	return &LikeDelivery{
		likeService: likeService,
	}
}

func (d *LikeDelivery) LikeRecipe(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	likeRequest := likes.LikeRequest{}
	err = e.Bind(&likeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	likeRequest.RecipeID = recipeId
	likeRequest.UserID = userId

	err = d.likeService.LikeRecipe(ConvertToEntity(&likeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.LIKE_SuccessLikeUserRecipe))
}

func (d *LikeDelivery) UnlikeRecipe(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	likeRequest := likes.LikeRequest{}
	err = e.Bind(&likeRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	likeRequest.RecipeID = recipeId
	likeRequest.UserID = userId

	err = d.likeService.UnlikeRecipe(ConvertToEntity(&likeRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.LIKE_SuccessUnlikeUserRecipe))
}
