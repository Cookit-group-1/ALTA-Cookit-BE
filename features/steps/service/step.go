package service

import (
	"alta-cookit-be/features/recipes"
	"alta-cookit-be/features/steps"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type StepSerivce struct {
	stepData steps.StepData_
	recipeData recipes.RecipeData_
	validate *validator.Validate
}

func New(stepData steps.StepData_, recipeData recipes.RecipeData_) steps.StepService_ {
	return &StepSerivce{
		stepData: stepData,
		recipeData: recipeData,
		validate: validator.New(),
	}
}

func (s *StepSerivce) InsertStep(entity *steps.StepEntity) (*steps.StepEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.recipeData.ActionValidator(entity.RecipeID, entity.UserID)
	if !isEntitled {
		return nil, errors.New(consts.SERVER_ForbiddenRequest)
	}

	output, err := s.stepData.InsertStep(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *StepSerivce) UpdateStepById(entity *steps.StepEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	isEntitled := s.stepData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err = s.stepData.UpdateStepById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *StepSerivce) DeleteStepById(entity *steps.StepEntity) error {
	isEntitled := s.stepData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.stepData.DeleteStepById(entity)
	if err != nil {
		return err
	}
	return nil
}
