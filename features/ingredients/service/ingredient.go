package service

import (
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type IngredientService struct {
	ingredientData ingredients.IngredientData_
	validate       *validator.Validate
}

func New(ingredientData ingredients.IngredientData_) ingredients.IngredientService_ {
	return &IngredientService{
		ingredientData: ingredientData,
		validate:       validator.New(),
	}
}

func (s *IngredientService) InsertIngredient(entity *ingredients.IngredientEntity) (*ingredients.IngredientEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
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

	err = s.ingredientData.UpdateIngredientById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *IngredientService) DeleteIngredientById(entity *ingredients.IngredientEntity) error {
	err := s.ingredientData.DeleteIngredientById(entity)
	if err != nil {
		return err
	}
	return nil
}
