package helpers

import (
	"alta-cookit-be/utils/consts"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Response(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}

func ResponseWithData(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

func ReturnBadResponse(e echo.Context, err error) error {
	switch err.Error() {
	case consts.JWT_InvalidJwtToken:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.JWT_FailedCastingJwtToken:
		return e.JSON(http.StatusInternalServerError, Response(err.Error()))

	case consts.ECHO_ErrorBindData:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.ECHO_InvaildIdParam:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.ECHO_InvaildPageParam:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.ECHO_InvaildLimitParam:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.ECHO_InvalidImageFileType:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))
	
	case consts.ECHO_InvalidFileSize:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.VALIDATION_InvalidInput:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.FOLLOWER_AlreadyFollowing:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.INGREDIENT_InvalidIngredient:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.LIKE_AlreadyLiked:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.RECIPE_InvalidRecipe:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.USER_InvalidUser:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case gorm.ErrRecordNotFound.Error():
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.SERVER_ForbiddenRequest:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	case consts.SERVER_InternalServerError:
		return e.JSON(http.StatusInternalServerError, Response(err.Error()))
		
	case consts.AUTH_SecurePassword:
		return e.JSON(http.StatusBadRequest, Response(err.Error()))

	default:
		return e.JSON(http.StatusInternalServerError, Response(err.Error()))
	}
}

func ErrorResponse(err error) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	msg := err.Error()

	if msg != "" {
		resp["message"] = msg
	}

	switch true {

	// error response for auth features
	case strings.Contains(msg, "Atoi"):
		resp["message"] = "id must be number, cannot be string"
		code = http.StatusNotFound
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_NotFound):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validate"):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.AUTH_ErrorCreateToken):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorHash):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorComparePassword):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.QUERY_ErrorInsertData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_NoRowsAffected):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorRole):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.QUERY_ErrorUpdateData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_ErrorEmptyPassword):
		code = http.StatusBadRequest
	case strings.Contains(msg, consts.QUERY_ErrorDeleteData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.QUERY_ErrorReadData):
		code = http.StatusInternalServerError
	case strings.Contains(msg, consts.AUTH_SecurePassword):
		code = http.StatusBadRequest
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "Unauthorized") {
		code = http.StatusUnauthorized
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	} else if strings.Contains(msg, "secure_password") {
		log.Println("error running register service: the password does not meet security requirements")
		code = http.StatusBadRequest
		resp["message"] = "password must be at least 8 characters long, must contain uppercase letters, must contain lowercase letters, must contain numbers, must not be too general"

	} else if strings.Contains(msg, "already exist") {
		words := strings.Split(msg, ": ")
		log.Println("error running " + words[0] + " service: already exist")
		resp["message"] = words[0] + " already exist"
		code = http.StatusConflict

	} else {
		log.Println("error running register service: required fields")
		code = http.StatusBadRequest
		resp["message"] = "required fields must be filled"
	}

	return code, resp
}
