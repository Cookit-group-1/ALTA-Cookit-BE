package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uint
	RecipeID uint
	Comment  string `gorm:"not null;type:text"`
	UrlImage string `gorm:"not null;type:text;default:''"`
}
