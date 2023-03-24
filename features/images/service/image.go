package service

import (
	"alta-cookit-be/features/images"

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
	output, err := s.imageData.UpdateImageById(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *ImageService) DeleteImageById(entity *images.ImageEntity) error {
	err := s.imageData.DeleteImageById(entity)
	if err != nil {
		return err
	}
	return nil
}
