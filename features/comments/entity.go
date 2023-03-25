package comments

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CommentEntity struct {
	ID             uint
	UserID         uint `validate:"required"`
	UserName       string
	UserRole       string
	ProfilePicture string
	RecipeID       uint   `validate:"required"`
	Comment        string `validate:"required"`
	Image          multipart.File
	ImageName      string
	UrlImage       string
}

type CommentRequest struct {
	ID        uint
	UserID    uint `form:"user_id"`
	UserRole  string
	RecipeID  uint
	Comment   string `form:"comment"`
	Image     multipart.File
	ImageName string
}

type CommentResponse struct {
	ID             uint   `json:"id,omitempty"`
	UserName       string `json:"username,omitempty"`
	UserRole       string `json:"user_role,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Comment        string `json:"comment,omitempty"`
	UrlImage       string `json:"comment_image,omitempty"`
}

type CommentDelivery_ interface {
	SelectCommentsByRecipeId(e echo.Context) error
	InsertComment(e echo.Context) error
	UpdateCommentById(e echo.Context) error
	DeleteCommentById(e echo.Context) error
}

type CommentService_ interface {
	SelectCommentsByRecipeId(commentEntity *CommentEntity) (*[]CommentEntity, error)
	InsertComment(commentEntity *CommentEntity) (*CommentEntity, error)
	UpdateCommentById(commentEntity *CommentEntity) (*CommentEntity, error)
	DeleteCommentById(commentEntity *CommentEntity) error
}

type CommentData_ interface {
	SelectCommentsByRecipeId(commentEntity *CommentEntity) (*[]CommentEntity, error)
	InsertComment(commentEntity *CommentEntity) (*CommentEntity, error)
	UpdateCommentById(commentEntity *CommentEntity) (*CommentEntity, error)
	DeleteCommentById(commentEntity *CommentEntity) error
}
