package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	RecipeID uint
	UrlImage string `gorm:"default:'';not null;type:VARCHAR(1000)"`
}
