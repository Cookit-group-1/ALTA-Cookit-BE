package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_ingredientData "alta-cookit-be/features/ingredients/data"
	_ingredientDelivery "alta-cookit-be/features/ingredients/delivery"
	_ingredientService "alta-cookit-be/features/ingredients/service"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	// userData := _userData.New(db)
	// userService := _userService.New(userData)
	// userHandler := _userDelivery.New(userService)

	// e.POST("/login", userHandler.Login)
	// e.POST("/users", userHandler.Register)
	// e.GET("/users", userHandler.GetUserData, middlewares.JWTMiddleware())
	// e.PUT("/users", userHandler.UpdateAccount, middlewares.JWTMiddleware())
	// e.PUT("/users/password", userHandler.UpdatePassword, middlewares.JWTMiddleware())
	// e.DELETE("/users", userHandler.RemoveAccount, middlewares.JWTMiddleware())
	// e.GET("/users/balances", userHandler.GetUserBalance, middlewares.JWTMiddleware())
	// e.PUT("/users/balances", userHandler.UpdateBalance, middlewares.JWTMiddleware())
}

func initIngredientRouter(db *gorm.DB, e *echo.Echo) {
	data := _ingredientData.New(db)
	service := _ingredientService.New(data)
	handler := _ingredientDelivery.New(service)

	e.POST("/ingredients", handler.InsertIngredient)
	e.PUT("/ingredients/:id", handler.UpdateIngredientById)
	e.DELETE("/ingredients/:id", handler.DeleteIngredientById)
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initIngredientRouter(db, e)
}
