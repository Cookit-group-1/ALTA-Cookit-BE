package likes

import (
	"github.com/labstack/echo/v4"
)

type LikeEntity struct {
	ID       uint
	UserID   uint
	RecipeID uint
}

type LikeRequest struct {
	ID       uint
	UserID   uint
	RecipeID uint
}

type LikeDelivery_ interface {
	LikeRecipe(e echo.Context) error
	UnlikeRecipe(e echo.Context) error
}

type LikeService_ interface {
	LikeRecipe(entity *LikeEntity) error
	UnlikeRecipe(entity *LikeEntity) error
}

type LikeData_ interface {
	LikeRecipe(entity *LikeEntity) error
	UnlikeRecipe(entity *LikeEntity) error
}
