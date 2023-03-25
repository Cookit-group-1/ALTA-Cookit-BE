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
		userID, _, _ := middlewares.ExtractToken(c)
		err := uh.srv.Deactive(userID)
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helpers.Response(consts.AUTH_ErrorBind))
	}
}

// Profile implements users.UserHandler
func (uh *userHandler) Profile() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements users.UserHandler
func (uh *userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
