package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uint
	RecipeID uint
	Comment  string `gorm:"default:'';not null;type:VARCHAR(1000)"`
	UrlImage string `gorm:"default:'';not null;type:VARCHAR(1000);"`
}
