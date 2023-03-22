package models

import (
	_cartModel "alta-cookit-be/features/carts/models"
	_commentModel "alta-cookit-be/features/comments/models"
	_followerModel "alta-cookit-be/features/followers/models"
	_likeModel "alta-cookit-be/features/likes/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
	_transactionModel "alta-cookit-be/features/transactions/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ProfilePicture string                          `gorm:"type:text"`
	Name           string                          `gorm:"not null;type:VARCHAR(50)"`
	Username       string                          `gorm:"unique;not null;type:VARCHAR(50)"`
	Role           string                          `gorm:"not null;type:enum('Admin', 'User', 'VerifiedUser');default:'User'"`
	Email          string                          `gorm:"unique;not null;type:VARCHAR(100)"`
	GoogleId       string                          `gorm:"unique;type:text"`
	Password       string                          `gorm:"not null;type:text"`
	Balance        float64                         `gorm:"not null;default:2000000"`
	FromFollowers  []_followerModel.Follower       `gorm:"foreignKey:FromUserID;constraint:OnDelete:CASCADE;"`
	ToFollowers    []_followerModel.Follower       `gorm:"foreignKey:ToUserID;constraint:OnDelete:CASCADE;"`
	Recipe         []_recipeModel.Recipe           `gorm:"constraint:OnDelete:CASCADE;"`
	Likes          []_likeModel.Like               `gorm:"constraint:OnDelete:CASCADE;"`
	Comments       []_commentModel.Comment         `gorm:"constraint:OnDelete:CASCADE;"`
	Carts          []_cartModel.Cart               `gorm:"constraint:OnDelete:CASCADE;"`
	Transactions   []_transactionModel.Transaction `gorm:"constraint:OnDelete:CASCADE;"`
}
