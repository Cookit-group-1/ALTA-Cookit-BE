package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_likeData "alta-cookit-be/features/likes/data"
	_likeDelivery "alta-cookit-be/features/likes/delivery"
	_likeService "alta-cookit-be/features/likes/service"
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
	_recipeData "alta-cookit-be/features/recipes/data"
	_recipeDelivery "alta-cookit-be/features/recipes/delivery"
	_recipeService "alta-cookit-be/features/recipes/service"
	_stepData "alta-cookit-be/features/steps/data"
	_stepDelivery "alta-cookit-be/features/steps/delivery"
	_stepService "alta-cookit-be/features/steps/service"
	_userData "alta-cookit-be/features/users/data"
	_userDelivery "alta-cookit-be/features/users/handler"
	_userService "alta-cookit-be/features/users/services"
	"alta-cookit-be/middlewares"
	"alta-cookit-be/utils/consts"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userDelivery.New(userService)

	// Auth
	e.POST("/login", userHandler.Login())
	e.POST("/register", userHandler.Register())

	// Users
	e.GET("/users", userHandler.Profile(), middlewares.JWTMiddleware())
	e.PUT("/users", userHandler.Update(), middlewares.JWTMiddleware())
	e.DELETE("/users", userHandler.Deactive(), middlewares.JWTMiddleware())
	e.POST("users/upgrade", userHandler.UpgradeUser(), middlewares.JWTMiddleware())
	e.GET("users/search", userHandler.SearchUser(), middlewares.JWTMiddleware())
	e.PUT("/users/password", userHandler.UpdatePassword(), middlewares.JWTMiddleware())
	e.GET("/users/:id", userHandler.ShowAnotherUserByID(), middlewares.JWTMiddleware())
	// e.GET("/users/balances", userHandler.GetUserBalance, middlewares.JWTMiddleware())
	// e.PUT("/users/balances", userHandler.UpdateBalance, middlewares.JWTMiddleware())

	// Admin
	e.PUT("/users/approval/:id", userHandler.AdminApproval(), middlewares.JWTMiddleware())
	e.GET("users/verifying", userHandler.ListUserRequest(), middlewares.JWTMiddleware())
}

func initRecipeRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	imageData := _imageData.New(db)
	data := _recipeData.New(db, userData, imageData)
	service := _recipeService.New(data)
	handler := _recipeDelivery.New(service)

	e.GET("/recipes", handler.SelectRecipes)
	e.POST("/recipes", handler.InsertRecipe, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s", consts.ECHO_P_RecipeId), handler.UpdateRecipeById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s", consts.ECHO_P_RecipeId), handler.DeleteRecipeById, middlewares.JWTMiddleware())
	e.GET("users/recipes/timeline", handler.SelectRecipesTimeline, middlewares.JWTMiddleware())
	e.GET("/recipes/trending", handler.SelectRecipesTrending)
	e.GET(fmt.Sprintf("/recipes/:%s/detail", consts.ECHO_P_RecipeId), handler.SelectRecipeDetailById)
}

func initImageRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	imageData := _imageData.New(db)
	recipeData := _recipeData.New(db, userData, imageData)
	data := _imageData.New(db)
	service := _imageService.New(data, recipeData)
	handler := _imageDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/images", consts.ECHO_P_RecipeId), handler.InsertImage, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s/images/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_ImageId), handler.UpdateImageById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s/images/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_ImageId), handler.DeleteImageById, middlewares.JWTMiddleware())
}

func initLikeRouter(db *gorm.DB, e *echo.Echo) {
	data := _likeData.New(db)
	service := _likeService.New(data)
	handler := _likeDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/like", consts.ECHO_P_RecipeId), handler.LikeRecipe, middlewares.JWTMiddleware())
	e.POST(fmt.Sprintf("/recipes/:%s/unlike", consts.ECHO_P_RecipeId), handler.UnlikeRecipe, middlewares.JWTMiddleware())
}

func initCommentRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	data := _commentData.New(db, userData)
	service := _commentService.New(data)
	handler := _commentDelivery.New(service)

	e.GET(fmt.Sprintf("/recipes/:%s/comments", consts.ECHO_P_RecipeId), handler.SelectCommentsByRecipeId)
	e.POST(fmt.Sprintf("/recipes/:%s/comments", consts.ECHO_P_RecipeId), handler.InsertComment, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s/comments/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_CommentId), handler.UpdateCommentById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s/comments/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_CommentId), handler.DeleteCommentById, middlewares.JWTMiddleware())
}

func initStepRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	imageData := _imageData.New(db)
	recipeData := _recipeData.New(db, userData, imageData)
	data := _stepData.New(db)
	service := _stepService.New(data, recipeData)
	handler := _stepDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/steps", consts.ECHO_P_RecipeId), handler.InsertStep, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s/steps/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_StepId), handler.UpdateStepById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s/steps/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_StepId), handler.DeleteStepById, middlewares.JWTMiddleware())
}

func initIngredientRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	imageData := _imageData.New(db)
	recipeData := _recipeData.New(db, userData, imageData)
	data := _ingredientData.New(db)
	service := _ingredientService.New(data, recipeData)
	handler := _ingredientDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/ingredients", consts.ECHO_P_RecipeId), handler.InsertIngredient, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s/ingredients/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientId), handler.UpdateIngredientById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s/ingredients/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientId), handler.DeleteIngredientById, middlewares.JWTMiddleware())
}

func initIngredientDetailRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	imageData := _imageData.New(db)
	recipeData := _recipeData.New(db, userData, imageData)
	data := _ingredientDetailData.New(db)
	service := _ingredientDetailService.New(data, recipeData)
	handler := _ingredientDetailDelivery.New(service)

	e.POST(fmt.Sprintf("/recipes/:%s/ingredients/:%s/ingredientDetails", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientId), handler.InsertIngredientDetail, middlewares.JWTMiddleware())
	e.PUT(fmt.Sprintf("/recipes/:%s/ingredients/ingredientDetails/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientDetailId), handler.UpdateIngredientDetailById, middlewares.JWTMiddleware())
	e.DELETE(fmt.Sprintf("/recipes/:%s/ingredients/ingredientDetails/:%s", consts.ECHO_P_RecipeId, consts.ECHO_P_IngredientDetailId), handler.DeleteIngredientDetailById, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initRecipeRouter(db, e)
	initImageRouter(db, e)
	initLikeRouter(db, e)
	initCommentRouter(db, e)
	initStepRouter(db, e)
	initIngredientRouter(db, e)
	initIngredientDetailRouter(db, e)
	initUserRouter(db, e)
}
