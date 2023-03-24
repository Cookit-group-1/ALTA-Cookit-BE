package service

import (
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"
	"errors"

	"github.com/go-playground/validator"
)

type RecipeService struct {
	recipeData recipes.RecipeData_
	validate   *validator.Validate
}

func New(recipeData recipes.RecipeData_) recipes.RecipeService_ {
	return &RecipeService{
		recipeData: recipeData,
		validate:   validator.New(),
	}
}

func (s *RecipeService) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	output, err := s.recipeData.InsertRecipe(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}
