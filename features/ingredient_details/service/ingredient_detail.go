package service

import (
	"alta-cookit-be/features/ingredient_details"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type IngredientDetailService struct {
	recipeData           recipes.RecipeData_
	ingredientDetailData ingredient_details.IngredientDetailData_
	validate             *validator.Validate
}

func New(ingredientDetailData ingredient_details.IngredientDetailData_, recipeData recipes.RecipeData_) ingredient_details.IngredientDetailService_ {
	return &IngredientDetailService{
		recipeData:           recipeData,
		ingredientDetailData: ingredientDetailData,
		validate:             validator.New(),
	}
}

func (s *IngredientDetailService) InsertIngredientDetail(entity *ingredient_details.IngredientDetailEntity) (*ingredient_details.IngredientDetailEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, err
	}

	isEntitled := s.recipeData.ActionValidator(entity.RecipeID, entity.UserID)
	if !isEntitled {
		return nil, errors.New(consts.SERVER_ForbiddenRequest)
	}

	output, err := s.ingredientDetailData.InsertIngredientDetail(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *IngredientDetailService) UpdateIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.ingredientDetailData.ActionValidator(entity.ID, entity.IngredientID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err = s.ingredientDetailData.UpdateIngredientDetailById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *IngredientDetailService) DeleteIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	isEntitled := s.ingredientDetailData.ActionValidator(entity.ID, entity.IngredientID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.ingredientDetailData.DeleteIngredientDetailById(entity)
	if err != nil {
		return err
	}
	return nil
}
