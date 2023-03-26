package delivery

import (
	"alta-cookit-be/features/comments"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentDelivery struct {
	commentService comments.CommentService_
}

func New(commentService comments.CommentService_) comments.CommentDelivery_ {
	return &CommentDelivery{
		commentService: commentService,
	}
}

func (d *CommentDelivery) SelectCommentsByRecipeId(e echo.Context) error {
	page, limit := helpers.ExtractPageLimit(e)
	limit, offset := helpers.LimitOffsetConvert(page, limit)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	commentRequest := comments.CommentRequest{}
	err = e.Bind(&commentRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	commentRequest.RecipeID = recipeId
	commentRequest.DataLimit = limit
	commentRequest.DataOffset = offset

	output, err := d.commentService.SelectCommentsByRecipeId(ConvertToEntity(&commentRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.COMMENT_SuccessReadRecipeCommentList, ConvertToResponses(output)))
}

func (d *CommentDelivery) InsertComment(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	commentRequest := comments.CommentRequest{}
	err = e.Bind(&commentRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}

	file, fileName, _ := helpers.ExtractFile(e, "image")
	commentRequest.RecipeID = recipeId
	commentRequest.UserID = userId
	commentRequest.Image = file
	commentRequest.ImageName = fileName

	output, err := d.commentService.InsertComment(ConvertToEntity(&commentRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.COMMENT_SuccessInsertRecipeComment, ConvertToResponse(output)))
}

func (d *CommentDelivery) UpdateCommentById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_CommentId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	commentRequest := comments.CommentRequest{}
	err = e.Bind(&commentRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}

	file, fileName, _ := helpers.ExtractFile(e, "image")
	commentRequest.ID = id
	commentRequest.UserID = userId
	commentRequest.RecipeID = recipeId
	commentRequest.Image = file
	commentRequest.ImageName = fileName

	output, err := d.commentService.UpdateCommentById(ConvertToEntity(&commentRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusCreated, helpers.ResponseWithData(consts.COMMENT_SuccessUpdateRecipeComment, ConvertToResponse(output)))
}

func (d *CommentDelivery) DeleteCommentById(e echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(e)
	id, err := helpers.ExtractIDParam(e, consts.ECHO_P_CommentId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}
	recipeId, err := helpers.ExtractIDParam(e, consts.ECHO_P_RecipeId)
	if err != nil {
		return errors.New(consts.ECHO_InvaildIdParam)
	}

	commentRequest := comments.CommentRequest{}
	err = e.Bind(&commentRequest)
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	commentRequest.ID = id
	commentRequest.RecipeID = recipeId
	commentRequest.UserID = userId

	err = d.commentService.DeleteCommentById(ConvertToEntity(&commentRequest))
	if err != nil {
		return helpers.ReturnBadResponse(e, err)
	}
	return e.JSON(http.StatusOK, helpers.Response(consts.COMMENT_SuccessDeleteRecipeComment))
}