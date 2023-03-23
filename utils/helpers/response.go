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

func ReturnBadResponse(e echo.Context, err error) error {
	switch err.Error() {
		case consts.JWT_InvalidJwtToken:
			return e.JSON(http.StatusBadRequest, Response(err.Error()))

		case consts.JWT_FailedCastingJwtToken:
			return e.JSON(http.StatusInternalServerError, Response(err.Error()))

		case consts.ECHO_ErrorBindData:
			return e.JSON(http.StatusInternalServerError, Response(err.Error()))

		case consts.ECHO_InvaildIdParam:
			return e.JSON(http.StatusBadRequest, Response(err.Error()))

		case consts.ECHO_InvaildPageParam:
			return e.JSON(http.StatusBadRequest, Response(err.Error()))

		case consts.ECHO_InvaildLimitParam:
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
		
		default:
			return e.JSON(http.StatusInternalServerError, Response(err.Error()))
	}
}