package transaction_details

import (
	"alta-cookit-be/features/images"

	"github.com/labstack/echo/v4"
)

type TransactionDetailEntity struct {
	ID                  uint
	LoggedInUserID      uint
	SellerUserID        uint
	SellerUsername      string
	IngredientID        uint `validate:"required"`
	RecipeImageEntities []images.ImageEntity
	RecipeName          string
	IngredientName      string
	Price               float64
	Quantity            int `validate:"required"`
}

type TransactionDetailRequest struct {
	ID             uint `json:"-" form:"-"`
	LoggedInUserID uint `json:"-" form:"-"`
	IngredientID   uint `json:"ingredient_id" form:"ingredient_id"`
	Quantity       int  `json:"quantity" form:"quantity"`
}

type TransactionDetailResponse struct {
	ID                   uint                   `json:"id,omitempty"`
	SellerUserID         uint                   `json:"id_seller,omitempty"`
	SellerUsername       string                 `json:"seller_user_username,omitempty"`
	RecipeImageResponses []images.ImageResponse `json:"recipe_images,omitempty"`
	RecipeName           string                 `json:"recipe_name,omitempty"`
	IngredientName       string                 `json:"ingredient_name,omitempty"`
	Price                float64                `json:"price,omitempty"`
	Quantity             int                    `json:"quantity,omitempty"`
}

type TransactionDetailDelivery_ interface {
	SelectTransactionDetailById(e echo.Context) error
}

type TransactionDetailService_ interface {
	SelectTransactionDetailById(transactionEntity *TransactionDetailEntity) (*TransactionDetailEntity, error)
}

type TransactionDetailData_ interface {
	ActionValidator(id, loggedInUserId uint) bool
	SelectTransactionDetailById(transactionEntity *TransactionDetailEntity) (*TransactionDetailEntity, error)
}
