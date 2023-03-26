package handler

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv users.UserService
}

func New(srv users.UserService) users.UserHandler {
	return &userHandler{
		srv: srv,
	}
}

// Login implements users.UserHandler
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginReq{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.AUTH_ErrorBind))
		}

		token, res, errLogin := uh.srv.Login(input.Username, input.Password)
		if errLogin != nil {
			return c.JSON(helpers.ErrorResponse(errLogin))
		}
		dataResponse := map[string]any{
			"id":       res.ID,
			"username": res.Username,
			"role":     res.Role,
			"token":    token,
		}
		return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.AUTH_SuccessLogin, dataResponse))
	}
}

// Register implements users.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterReq{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.AUTH_ErrorBind))
		}

		res, errRegister := uh.srv.Register(*ReqToCore(input))
		if errRegister != nil {
			return c.JSON(helpers.ErrorResponse(errRegister))
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, helpers.Response(consts.AUTH_SuccessCreate))
	}
}

// Deactive implements users.UserHandler
func (uh *userHandler) Deactive() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		err := uh.srv.Deactive(id)
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helpers.Response(consts.AUTH_ErrorBind))
	}
}

// Profile implements users.UserHandler
func (uh *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		dataCore, err := uh.srv.Profile(id)
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.AUTH_ErrorBind, ToProfileResponse(dataCore)))
	}
}

// Update implements users.UserHandler
func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		input := UpdateProfileReq{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.AUTH_ErrorBind))
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("profile_picture")
		if checkFile != nil {
			formHeader, err := c.FormFile("profile_picture")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}

		res, err := uh.srv.Update(id, input.FileHeader, *ReqToCore(input))
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		result, err := ConvertUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "success update profile",
			})
		}
	}
}
