package models

import (
	_commentModel "alta-cookit-be/features/comments/models"
	_imageModel "alta-cookit-be/features/images/models"
	_ingredientModel "alta-cookit-be/features/ingredients/models"
	_likeModel "alta-cookit-be/features/likes/models"
	_stepModel "alta-cookit-be/features/steps/models"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	UserID      uint
	RecipeID    *uint
	Recipe      *Recipe
	Type        string                        `gorm:"not null;type:enum('Original', 'Mixed', 'Cooked');default:'Original'"`
	Status      string                        `gorm:"not null;type:enum('None', 'OpenForSale');default:'None'"`
	Name        string                        `gorm:"default:'';not null;type:VARCHAR(50)"`
	Description string                        `gorm:"default:'';not null;type:text"`
	Steps       []_stepModel.Step             `gorm:"constraint:OnDelete:CASCADE;"`
	Ingredients []_ingredientModel.Ingredient `gorm:"constraint:OnDelete:CASCADE;"`
	Comments    []_commentModel.Comment       `gorm:"constraint:OnDelete:CASCADE;"`
	Likes       []_likeModel.Like             `gorm:"constraint:OnDelete:CASCADE;"`
	Images      []_imageModel.Image           `gorm:"constraint:OnDelete:CASCADE;"`
}
