package service

import (
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"
	"errors"
)

type RecipeService struct {
	recipeData recipes.RecipeData_
}

func New(recipeData recipes.RecipeData_) recipes.RecipeService_ {
	return &RecipeService{
		recipeData: recipeData,
	}
}

func (s *RecipeService) InsertRecipe(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	output, err := s.recipeData.InsertRecipe(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *RecipeService) UpdateRecipeById(entity *recipes.RecipeEntity) error {
	isEntitled := s.recipeData.ActionValidator(entity.ID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.recipeData.UpdateRecipeById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) DeleteRecipeById(entity *recipes.RecipeEntity) error {
	isEntitled := s.recipeData.ActionValidator(entity.ID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.recipeData.DeleteRecipeById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) SelectRecipesByUserId(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	outputs, err := s.recipeData.SelectRecipesByUserId(entity)
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (s *RecipeService) SelectRecipesTimeline(entity *recipes.RecipeEntity) (*[]recipes.RecipeEntity, error) {
	output, err := s.recipeData.SelectRecipesTimeline(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *RecipeService) SelectRecipeDetailById(entity *recipes.RecipeEntity) (*recipes.RecipeEntity, error) {
	output, err := s.recipeData.SelectRecipeDetailById(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}
