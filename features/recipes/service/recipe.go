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

func (s *RecipeService) SelectRecipesByUserId(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	outputs, err := s.recipeData.SelectRecipesByUserId(entity)
	if err != nil {
		return nil, err
	}
	return outputs, nil
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

func (s *RecipeService) UpdateRecipeById(entity *recipes.RecipeEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	err = s.recipeData.UpdateRecipeById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) DeleteRecipeById(entity *recipes.RecipeEntity) error {
	err := s.recipeData.DeleteRecipeById(entity)
	if err != nil {
		return err
	}
	return nil
}
