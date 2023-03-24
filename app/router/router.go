package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_recipeData "alta-cookit-be/features/recipes/data"
	_recipeDelivery "alta-cookit-be/features/recipes/delivery"
	_recipeService "alta-cookit-be/features/recipes/service"
	_commentData "alta-cookit-be/features/comments/data"
	_commentDelivery "alta-cookit-be/features/comments/delivery"
	_commentService "alta-cookit-be/features/comments/service"
	_imageData "alta-cookit-be/features/images/data"
	_imageDelivery "alta-cookit-be/features/images/delivery"
	_imageService "alta-cookit-be/features/images/service"
	_ingredientDetailData "alta-cookit-be/features/ingredient_details/data"
	_ingredientDetailDelivery "alta-cookit-be/features/ingredient_details/delivery"
	_ingredientDetailService "alta-cookit-be/features/ingredient_details/service"
	_ingredientData "alta-cookit-be/features/ingredients/data"
	_ingredientDelivery "alta-cookit-be/features/ingredients/delivery"
	_ingredientService "alta-cookit-be/features/ingredients/service"
	_stepData "alta-cookit-be/features/steps/data"
	_stepDelivery "alta-cookit-be/features/steps/delivery"
	_stepService "alta-cookit-be/features/steps/service"
	_userData "alta-cookit-be/features/users/data"
	"alta-cookit-be/utils/consts"
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

func initRecipeRouter(db *gorm.DB, e *echo.Echo) {
	data := _recipeData.New(db)
	service := _recipeService.New(data)
	handler := _recipeDelivery.New(service)

	e.POST("/recipes", handler.InsertRecipe)
}

func initImageRouter(db *gorm.DB, e *echo.Echo) {
	data := _imageData.New(db)
	service := _imageService.New(data)
	handler := _imageDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/images", consts.ECHO_P_RecipeId), handler.InsertImage)
	e.PUT(fmt.Sprintf("/recipes/:%s/images/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_ImageId), handler.UpdateImageById)
	e.DELETE(fmt.Sprintf("/recipes/:%s/images/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_ImageId), handler.DeleteImageById)
}

func initCommentRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	data := _commentData.New(db, userData)
	service := _commentService.New(data)
	handler := _commentDelivery.New(service)

	e.GET(fmt.Sprintf("/recipes/:%s/comments", consts.ECHO_P_RecipeId), handler.SelectCommentsByRecipeId)
	e.POST(fmt.Sprintf("/recipes/:%s/comments", consts.ECHO_P_RecipeId), handler.InsertComment)
	e.PUT(fmt.Sprintf("/recipes/:%s/comments/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_CommentId), handler.UpdateCommentById)
	e.DELETE(fmt.Sprintf("/recipes/:%s/comments/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_CommentId), handler.DeleteCommentById)
}

func initStepRouter(db *gorm.DB, e *echo.Echo) {
	data := _stepData.New(db)
	service := _stepService.New(data)
	handler := _stepDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/steps", consts.ECHO_P_RecipeId), handler.InsertStep)
	e.PUT(fmt.Sprintf("/recipes/:%s/steps/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_StepId), handler.UpdateStepById)
	e.DELETE(fmt.Sprintf("/recipes/:%s/steps/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_StepId), handler.DeleteStepById)
}

func initIngredientRouter(db *gorm.DB, e *echo.Echo) {
	data := _ingredientData.New(db)
	service := _ingredientService.New(data)
	handler := _ingredientDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/ingredients", consts.ECHO_P_RecipeId), handler.InsertIngredient)
	e.PUT(fmt.Sprintf("/recipes/:%s/ingredients/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientId), handler.UpdateIngredientById)
	e.DELETE(fmt.Sprintf("/recipes/:%s/ingredients/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientId), handler.DeleteIngredientById)
}

func initIngredientDetailRouter(db *gorm.DB, e *echo.Echo) {
	data := _ingredientDetailData.New(db)
	service := _ingredientDetailService.New(data)
	handler := _ingredientDetailDelivery.New(service)

	e.POST(fmt.Sprintf("/ingredients/:%s/ingredientDetails", consts.ECHO_P_IngredientId), handler.InsertIngredientDetail)
	e.PUT(fmt.Sprintf("/ingredients/:%s/ingredientDetails/:%s", consts.ECHO_P_IngredientId, consts.ECHO_P_IngredientDetailId), handler.UpdateIngredientDetailById)
	e.DELETE(fmt.Sprintf("/ingredients/:%s/ingredientDetails/:%s", consts.ECHO_P_IngredientId, consts.ECHO_P_IngredientDetailId), handler.DeleteIngredientDetailById)
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initRecipeRouter(db, e)
	initImageRouter(db, e)
	initCommentRouter(db, e)
	initStepRouter(db, e)
	initIngredientRouter(db, e)
	initIngredientDetailRouter(db, e)
}
