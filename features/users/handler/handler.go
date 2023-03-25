package handler

import (
	"alta-cookit-be/features/users"
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
	panic("unimplemented")
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
