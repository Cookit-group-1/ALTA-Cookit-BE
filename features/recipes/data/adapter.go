package data

import (
	_imageData "alta-cookit-be/features/images/data"
	_ingredientData "alta-cookit-be/features/ingredients/data"
	"alta-cookit-be/features/recipes"
	_recipeModel "alta-cookit-be/features/recipes/models"
	_stepData "alta-cookit-be/features/steps/data"
	_userModel "alta-cookit-be/features/users/models"
)

func ConvertToGorm(entity *recipes.RecipeEntity) *_recipeModel.Recipe {
	gorm := _recipeModel.Recipe{
		UserID:      entity.UserID,
		RecipeID:    &entity.RecipeID,
		Type:        entity.Type,
		Status:      entity.Status,
		Name:        entity.Name,
		Description: entity.Description,
		Steps:       *_stepData.ConvertToGorms(&entity.StepEntities),
		Ingredients: *_ingredientData.ConvertToGorms(&entity.IngredientEntities),
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToEntity(gorm *_recipeModel.Recipe, userGorm ...*_userModel.User) *recipes.RecipeEntity {
	entity := recipes.RecipeEntity{
		ID:                 gorm.ID,
		UserID:             gorm.UserID,
		RecipeID:           *gorm.RecipeID,
		Type:               gorm.Type,
		Status:             gorm.Status,
		Name:               gorm.Name,
		Description:        gorm.Description,
		StepEntities:       *_stepData.ConvertToEntities(&gorm.Steps),
		IngredientEntities: *_ingredientData.ConvertToEntities(&gorm.Ingredients),
		ImageEntities:      *_imageData.ConvertToEntities(&gorm.Images),
	}
	if gorm.Recipe != nil {
		entity.Recipe = ConvertToEntity(gorm)
	}
	if len(userGorm) != 0 {
		entity.UserName = userGorm[0].Username
		entity.UserRole = userGorm[0].Role
		entity.ProfilePicture = userGorm[0].ProfilePicture
	}
	return &entity
}
