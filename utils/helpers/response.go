package helpers

import (
	"alta-cookit-be/utils/consts"
	"net/http"

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

func ValidateBadResponse(e echo.Context, err error) (codeStatus int, failedMessage string) {
	switch err.Error() {
		case consts.JWT_InvalidJwtToken:
			return http.StatusBadRequest, err.Error()

		case consts.JWT_FailedCastingJwtToken:
			return http.StatusInternalServerError, err.Error()

		case consts.ECHO_ErrorBindData:
			return http.StatusBadRequest, err.Error()

		case consts.ECHO_InvaildIdParam:
			return http.StatusBadRequest, err.Error()

		case consts.ECHO_InvaildPageParam:
			return http.StatusBadRequest, err.Error()

		case consts.ECHO_InvaildLimitParam:
			return http.StatusBadRequest, err.Error()

		case consts.VALIDATION_InvalidInput:
			return http.StatusBadRequest, err.Error()

		case consts.FOLLOWER_AlreadyFollowing:
			return http.StatusBadRequest, err.Error()

		case consts.INGREDIENT_InvalidIngredient:
			return http.StatusBadRequest, err.Error()

		case consts.LIKE_AlreadyLiked:
			return http.StatusBadRequest, err.Error()

		case consts.RECIPE_InvalidRecipe:
			return http.StatusBadRequest, err.Error()

		case consts.USER_InvalidUser:
			return http.StatusBadRequest, err.Error()

		case gorm.ErrRecordNotFound.Error():
			return http.StatusBadRequest, consts.GORM_RecordNotFound

		case consts.SERVER_ForbiddenRequest:
			return http.StatusBadRequest, err.Error()
			
		case consts.SERVER_InternalServerError:
			return http.StatusInternalServerError, err.Error()
		
		default:
			return http.StatusInternalServerError, err.Error()
	}
}