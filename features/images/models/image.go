package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	RecipeID uint
	UrlImage string `gorm:"not null;type:text"`
}
