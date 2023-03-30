package service

import (
	"alta-cookit-be/features/carts"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type CartService struct {
	cartData carts.CartData_
	validate *validator.Validate
}

func New(cartData carts.CartData_) carts.CartService_ {
	return &CartService{
		cartData: cartData,
		validate: validator.New(),
	}
}

func (s *CartService) SelectCartsByUserId(entity *carts.CartEntity) (*[]carts.CartEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	// isEntitled := s.cartData.ActionValidator(entity.RecipeID, entity.UserID)
	// if !isEntitled {
	// 	return nil, errors.New(consts.SERVER_ForbiddenRequest)
	// }

	outputs, err := s.cartData.SelectCartsByUserId(entity)
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (s *CartService) InsertCart(entity *carts.CartEntity) (*carts.CartEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	// isEntitled := s.cartData.ActionValidator(entity.RecipeID, entity.UserID)
	// if !isEntitled {
	// 	return nil, errors.New(consts.SERVER_ForbiddenRequest)
	// }

	output, err := s.cartData.InsertCart(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *CartService) UpdateCartById(entity *carts.CartEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	// isEntitled := s.ingredientData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	// if !isEntitled {
	// 	return errors.New(consts.SERVER_ForbiddenRequest)
	// }

	err = s.cartData.UpdateCartById(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) DeleteCartById(entity *carts.CartEntity) error {
	// isEntitled := s.ingredientData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	// if !isEntitled {
	// 	return errors.New(consts.SERVER_ForbiddenRequest)
	// }

	err := s.cartData.DeleteCartById(entity)
	if err != nil {
		return err
	}
	return nil
}
