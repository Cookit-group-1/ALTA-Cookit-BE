package service

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/utils/consts"
	"errors"

	"github.com/go-playground/validator"
)

type ImageService struct {
	imageData images.ImageData_
	validate  *validator.Validate
}

func New(imageData images.ImageData_) images.ImageService_ {
	return &ImageService{
		imageData: imageData,
		validate:  validator.New(),
	}
}

func (s *ImageService) InsertImage(entity *[]images.ImageEntity) (*[]images.ImageEntity, error) {
	output, err := s.imageData.InsertImage(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *ImageService) UpdateImageById(entity *images.ImageEntity) (*images.ImageEntity, error) {
	isEntitled := s.imageData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return nil, errors.New(consts.SERVER_ForbiddenRequest)
	}

	output, err := s.imageData.UpdateImageById(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *ImageService) DeleteImageById(entity *images.ImageEntity) error {
	isEntitled := s.imageData.ActionValidator(entity.ID, entity.RecipeID, entity.UserID)
	if !isEntitled {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := s.imageData.DeleteImageById(entity)
	if err != nil {
		return err
	}
	return nil
}
