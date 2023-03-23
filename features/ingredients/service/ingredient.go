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

func (s *IngredientService) InsertIngredient(ingredientEntity *ingredients.IngredientEntity) (*ingredients.IngredientEntity, error) {
	err := s.validate.Struct(ingredientEntity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	output, err := s.ingredientData.InsertIngredient(ingredientEntity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *IngredientService) UpdateIngredientById(ingredientEntity *ingredients.IngredientEntity) error {
	err := s.validate.Struct(ingredientEntity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	err = s.ingredientData.UpdateIngredientById(ingredientEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *IngredientService) DeleteIngredientById(ingredientEntity *ingredients.IngredientEntity) error {
	err := s.validate.Struct(ingredientEntity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	err = s.ingredientData.DeleteIngredientById(ingredientEntity)
	if err != nil {
		return err
	}
	return nil
}
