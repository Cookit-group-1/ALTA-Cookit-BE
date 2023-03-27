package images

import (
	"mime/multipart"

	_imageModel "alta-cookit-be/features/images/models"

	"github.com/labstack/echo/v4"
)

type ImageEntity struct {
	ID             uint
	UserID         uint 
	RecipeID       uint
	Image          multipart.File
	ImageName      string
	UrlImage       string
}

type ImageRequest struct {
	ID        uint
	UserID    uint
	RecipeID  uint
	Image     multipart.File
	ImageName string
}

type ImageResponse struct {
	ID             uint   `json:"id,omitempty"`
	UrlImage       string `json:"url_image,omitempty"`
}

type ImageDelivery_ interface {
	InsertImage(e echo.Context) error
	UpdateImageById(e echo.Context) error
	DeleteImageById(e echo.Context) error
}

type ImageService_ interface {
	InsertImage(imageEntity *[]ImageEntity) (*[]ImageEntity, error)
	UpdateImageById(imageEntity *ImageEntity) (*ImageEntity, error)
	DeleteImageById(imageEntity *ImageEntity) error
}

type ImageData_ interface {
	ActionValidator(id, recipeId, userId uint) bool
	SelectImagesByRecipeId(recipeId uint) *[]_imageModel.Image
	InsertImage(imageEntity *[]ImageEntity) (*[]ImageEntity, error)
	UpdateImageById(imageEntity *ImageEntity) (*ImageEntity, error)
	DeleteImageById(imageEntity *ImageEntity) error
}
