package recipes

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/features/steps"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type RecipeEntity struct {
	ID                   uint
	UserID               uint
	UserName             string
	UserRole             string
	ProfilePicture       string
	RecipeID             uint
	Image                []multipart.File
	ImageName            []string
	Recipe               *RecipeEntity
	UrlImage             string
	Type                 string
	Status               string
	Name                 string 
	Description          string 
	TotalLike            int    `gorm:"column:total_like"`
	TotalComment         int    `gorm:"column:total_comment"`
	StepEntities         []steps.StepEntity
	IngredientEntities   []ingredients.IngredientEntity
	ImageEntities        []images.ImageEntity
	DataLimit            int
	DataOffset           int
	ExtractedQueryParams map[string]interface{}
}

type RecipeRequest struct {
	ID                   uint
	UserID               uint
	UserRole             string
	RecipeID             uint `json:"recipe_id" form:"recipe_id"`
	Image                []multipart.File
	ImageName            []string
	Type                 string                          `json:"type" form:"type"`
	Status               string                          `json:"status" form:"status"`
	Name                 string                          `json:"name" form:"name"`
	Description          string                          `json:"description" form:"description"`
	StepRequests         []steps.StepRequest             `json:"steps" form:"steps"`
	IngredientRequests   []ingredients.IngredientRequest `json:"ingredients" form:"ingredients"`
	DataLimit            int
	DataOffset           int
	ExtractedQueryParams map[string]interface{}
}

type RecipeResponse struct {
	ID                  uint                             `json:"id,omitempty"`
	UserID              uint                             `json:"user_id,omitempty"`
	UserName            string                           `json:"username,omitempty"`
	UserRole            string                           `json:"user_role,omitempty"`
	ProfilePicture      string                           `json:"profile_picture,omitempty"`
	RecipeID            uint                             `json:"recipe_id,omitempty"`
	Recipe              *RecipeResponse                  `json:"replied_recipe,omitempty"`
	UrlImage            string                           `json:"url_image,omitempty"`
	Type                string                           `json:"type,omitempty"`
	Status              string                           `json:"status,omitempty"`
	Name                string                           `json:"name,omitempty"`
	Description         string                           `json:"description,omitempty"`
	TotalLike           int                              `json:"total_like,omitempty"`
	TotalComment        int                              `json:"total_comment,omitempty"`
	StepResponses       []steps.StepResponse             `json:"steps,omitempty"`
	IngredientResponses []ingredients.IngredientResponse `json:"ingredients,omitempty"`
	ImageResponses      []images.ImageResponse           `json:"images,omitempty"`
}

type RecipeDelivery_ interface {
	SelectRecipes(e echo.Context) error
	InsertRecipe(e echo.Context) error
	UpdateRecipeById(e echo.Context) error
	DeleteRecipeById(e echo.Context) error
	SelectRecipesTimeline(e echo.Context) error
	SelectRecipesTrending(e echo.Context) error
	SelectRecipeDetailById(e echo.Context) error
}

type RecipeService_ interface {
	SelectRecipes(entity *RecipeEntity) (*[]RecipeEntity, error)
	InsertRecipe(entity *RecipeEntity) (*RecipeEntity, error)
	UpdateRecipeById(entity *RecipeEntity) error
	DeleteRecipeById(entity *RecipeEntity) error
	SelectRecipesTimeline(entity *RecipeEntity) (*[]RecipeEntity, error)
	SelectRecipesTrending(entity *RecipeEntity) (*[]RecipeEntity, error)
	SelectRecipeDetailById(entity *RecipeEntity) (*RecipeEntity, error)
}

type RecipeData_ interface {
	ActionValidator(recipeId, userId uint) bool
	SelectRecipes(entity *RecipeEntity) (*[]RecipeEntity, error)
	InsertRecipe(entity *RecipeEntity) (*RecipeEntity, error)
	UpdateRecipeById(entity *RecipeEntity) error
	DeleteRecipeById(entity *RecipeEntity) error
	SelectRecipesTimeline(entity *RecipeEntity) (*[]RecipeEntity, error)
	SelectRecipesTrending(entity *RecipeEntity) (*[]RecipeEntity, error)
	SelectRecipeDetailById(entity *RecipeEntity) (*RecipeEntity, error)
}
