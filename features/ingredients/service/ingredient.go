package service

import (
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type IngredientService struct {
	recipeData     recipes.RecipeData_
	ingredientData ingredients.IngredientData_
	validate       *validator.Validate
}

func New(ingredientData ingredients.IngredientData_, recipeData recipes.RecipeData_) ingredients.IngredientService_ {
	return &IngredientService{
		recipeData:     recipeData,
		ingredientData: ingredientData,
		validate:       validator.New(),
	}
}

func (s *IngredientService) InsertIngredient(entity *ingredients.IngredientEntity) (*ingredients.IngredientEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.recipeData.ActionValidator(entity.RecipeID, entity.UserID)
	if !isEntitled {
		return nil, errors.New(consts.SERVER_ForbiddenRequest)
	}

	output, err := s.ingredientData.InsertIngredient(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *IngredientService) UpdateIngredientById(entity *ingredients.IngredientEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.ingredientData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err = s.ingredientData.UpdateIngredientById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *IngredientService) DeleteIngredientById(entity *ingredients.IngredientEntity) error {
	isEntitled := s.ingredientData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.ingredientData.DeleteIngredientById(entity)
	if err != nil {
		return err
	}
	return nil
}
