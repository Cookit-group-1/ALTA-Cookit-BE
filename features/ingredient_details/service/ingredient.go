package service

import (
	"alta-cookit-be/features/ingredient_details"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type IngredientDetailService struct {
	ingredientDetailData ingredient_details.IngredientDetailData_
	validate       *validator.Validate
}

func New(ingredientDetailData ingredient_details.IngredientDetailData_) ingredient_details.IngredientDetailService_ {
	return &IngredientDetailService{
		ingredientDetailData: ingredientDetailData,
		validate:       validator.New(),
	}
}

func (s *IngredientDetailService) InsertIngredientDetail(entity *ingredient_details.IngredientDetailEntity) (*ingredient_details.IngredientDetailEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
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

	err = s.ingredientDetailData.UpdateIngredientDetailById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *IngredientDetailService) DeleteIngredientDetailById(entity *ingredient_details.IngredientDetailEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	err = s.ingredientDetailData.DeleteIngredientDetailById(entity)
	if err != nil {
		return err
	}
	return nil
}
