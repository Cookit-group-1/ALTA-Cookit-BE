package handler

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
	"alta-cookit-be/utils/helpers"
	"log"
	"net/http"
	"strconv"
	"strings"

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
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.ECHO_ErrorBindData))
		}

		res, errRegister := uh.srv.Register(*ReqToCore(input))
		if errRegister != nil {
			return c.JSON(helpers.PrintErrorResponse(errRegister.Error()))
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
		return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessDelete))
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
		return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_SuccessGetProfile, ToProfileResponse(dataCore)))
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

// UpdatePassword implements users.UserHandler
func (uh *userHandler) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _, _ := middlewares.ExtractToken(c)
		input := UpdatePasswordReq{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.AUTH_ErrorBind))
		}

		err := uh.srv.UpdatePassword(userID, *ReqToCore(input))
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessUpdatePassword))
	}
}

// UpgradeUser implements users.UserHandler
func (uh *userHandler) UpgradeUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role, _ := middlewares.ExtractToken(c)
		if role == "VerifiedUser" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "your role already VerifiedUser"})
		}
		input := ApprovementReq{
			Approvement: "requested",
		}

		_, err := uh.srv.UpgradeUser(userID, *ReqToCore(input))

		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success send your request to admin",
		})
	}
}

// Search implements users.UserHandler
func (uh *userHandler) SearchUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _, _ := middlewares.ExtractToken(c)
		quotes := c.QueryParam("q")
		log.Println(quotes)
		res, err := uh.srv.SearchUser(userID, quotes)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		if quotes == "Admin" || quotes == "admin" || quotes == "ADMIN" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "cannot search credentials data"})
		}
		if quotes == "" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		result := []SearchResponse{}
		for i := 0; i < len(res); i++ {
			result = append(result, ToSearchResponse(res[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success find user",
		})
	}
}

// ShowAnotherUserByID implements users.UserHandler
func (uh *userHandler) ShowAnotherUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		pID := c.Param("id")
		anotherUserID, _ := strconv.Atoi(pID)
		_, _, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "error from server"})
		}
		dataCore, err := uh.srv.Profile(uint(anotherUserID))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_SuccessGetProfile, ToProfileResponse(dataCore)))
	}
}

// AdminApproval implements users.UserHandler
func (uh *userHandler) AdminApproval() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, _ := middlewares.ExtractToken(c)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "only admin can pass"})
		}
		approveID := c.Param("id")
		approvementID, _ := strconv.Atoi(approveID)

		input := ApprovalReq{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "cannot verify user, please choose verify or deny"})
		}

		targetRole := "User"
		if input.Status == "verify" {
			targetRole = "VerifiedUser"
		} else {
			return c.JSON(http.StatusCreated, map[string]interface{}{"message": "user request verified has been denied"})
		}

		updatedApprovement := ApprovementReq{
			Approvement: input.Status,
			Role:        targetRole,
		}

		_, err := uh.srv.UpgradeUser(uint(approvementID), *ReqToCore(updatedApprovement))

		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success approve verified user request",
		})
	}
}

// ListUserRequest implements users.UserHandler
func (uh *userHandler) ListUserRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role, _ := middlewares.ExtractToken(c)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "only admin can pass"})
		}
		dataCore, err := uh.srv.ListUserRequest(id)
		if err != nil {
			return c.JSON(helpers.ErrorResponse(err))
		}
		result := []ListUserRequestedResponse{}
		for i := 0; i < len(dataCore); i++ {
			result = append(result, ToListUserRequestedResponse(dataCore[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all requested users",
		})
	}
}
