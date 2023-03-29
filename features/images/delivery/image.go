package delivery

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageDelivery struct {
	imageService images.ImageService_
}

func New(imageService images.ImageService_) images.ImageDelivery_ {
	return &ImageDelivery{
		imageService: imageService,
	}
}

func (d *ImageDelivery) InsertImage(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	imageRequests := []images.ImageRequest{}
	err = e.Bind(&imageRequests)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}

	files, fileNames, err := helpers.ExtractMultipleImageFiles(e, "image")
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	for index, file := range files {
		imageRequests = append(imageRequests, images.ImageRequest{
			UserID: userId,
			RecipeID: recipeId,
			Image: file,
			ImageName: fileNames[index],
		})
	}

	output, err := d.imageService.InsertImage(ConvertToEntities(&imageRequests))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.IMAGE_SuccessInsertRecipeImage, ConvertToResponses(output)))
}

func (d *ImageDelivery) UpdateImageById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_ImageId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	imageRequest := images.ImageRequest{}
	err = e.Bind(&imageRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}

	file, fileName, _ := helpers.ExtractImageFile(e, "image")
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	imageRequest.ID = id
	imageRequest.UserID = userId
	imageRequest.RecipeID = recipeId
	imageRequest.Image = file
	imageRequest.ImageName = fileName

	output, err := d.imageService.UpdateImageById(ConvertToEntity(&imageRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.IMAGE_SuccesUpdateRecipeImage, ConvertToResponse(output)))
}

func (d *ImageDelivery) DeleteImageById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_CommentId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	imageRequest := images.ImageRequest{}
	err = e.Bind(&imageRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	imageRequest.ID = id
	imageRequest.RecipeID = recipeId
	imageRequest.UserID = userId

	err = d.imageService.DeleteImageById(ConvertToEntity(&imageRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.IMAGE_SuccessDeleteRecipeImage))
}