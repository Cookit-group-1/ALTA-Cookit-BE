package data

import (
	"alta-cookit-be/features/images"
	_imageModel "alta-cookit-be/features/images/models"
)

func ConvertToGorm(entity *images.ImageEntity) *_imageModel.Image {
	gorm := _imageModel.Image{
		RecipeID: entity.RecipeID,
		UrlImage: entity.UrlImage,
	}
	if entity.ID != 0 {
		gorm.ID = entity.ID
	}
	return &gorm
}

func ConvertToGorms(entities *[]images.ImageEntity) *[]_imageModel.Image {
	gorms := []_imageModel.Image{}
	for _, entity := range *entities {
		gorms = append(gorms, *ConvertToGorm(&entity))
	}
	return &gorms
}

func ConvertToEntity(gorm *_imageModel.Image) *images.ImageEntity {
	entity := images.ImageEntity{
		ID:       gorm.ID,
		RecipeID: gorm.RecipeID,
		UrlImage: gorm.UrlImage,
	}
	return &entity
}

func ConvertToEntities(gorms *[]_imageModel.Image) *[]images.ImageEntity {
	entities := []images.ImageEntity{}
	for _, gorm := range *gorms {
		entities = append(entities, *ConvertToEntity(&gorm))
	}
	return &entities
}
