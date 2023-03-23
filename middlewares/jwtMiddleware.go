package middlewares

import (
	"alta-cookit-be/app/config"
	"alta-cookit-be/utils/consts"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims[consts.JWT_Authorized] = true
	claims[consts.JWT_UserId] = userId
	claims[consts.JWT_Role] = userRole
	claims[consts.JWT_ExpiredTime] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(e echo.Context) (uint, string, error) {
	token, ok := e.Get("user").(*jwt.Token)
	if !ok {
		return 0, "", errors.New(consts.JWT_InvalidJwtToken)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New(consts.JWT_FailedCastingJwtToken)
	}
	if token.Valid {
		loggedInUserId, existedUserId := claims[consts.JWT_UserId].(float64)
		loggedInUserRole, existedUserRole := claims[consts.JWT_Role]
		if !existedUserId || !existedUserRole {
			return 0, "", errors.New(consts.SERVER_InternalServerError)
		}
		return uint(loggedInUserId), loggedInUserRole.(string), nil
	}
	return 0, "", errors.New(consts.SERVER_InternalServerError)
}