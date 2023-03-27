package data

import (
	_cartModel "alta-cookit-be/features/carts/models"
	_commentModel "alta-cookit-be/features/comments/models"
	_followerModel "alta-cookit-be/features/followers/models"
	_likeModel "alta-cookit-be/features/likes/models"
	_recipeModel "alta-cookit-be/features/recipes/models"
	_transactionModel "alta-cookit-be/features/transactions/models"
	"alta-cookit-be/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ProfilePicture string `gorm:"type:VARCHAR(1000)"`
	Username       string `gorm:"unique;not null;type:VARCHAR(50)"`
	Bio            string `gorm:"not null;type:VARCHAR(151)"`
	Role           string `gorm:"not null;type:enum('Admin', 'User', 'VerifiedUser');default:'User'"`
	Email          string `gorm:"unique;not null;type:VARCHAR(100)"`
	GoogleId       string `gorm:"type:VARCHAR(1000)"`
	Password       string `gorm:"not null;type:VARCHAR(1000)"`
	Approvement    string //no, requested and accepted
	FromFollowers  []_followerModel.Follower       `gorm:"foreignKey:FromUserID;constraint:OnDelete:CASCADE;"`
	ToFollowers    []_followerModel.Follower       `gorm:"foreignKey:ToUserID;constraint:OnDelete:CASCADE;"`
	Recipe         []_recipeModel.Recipe           `gorm:"constraint:OnDelete:CASCADE;"`
	Likes          []_likeModel.Like               `gorm:"constraint:OnDelete:CASCADE;"`
	Comments       []_commentModel.Comment         `gorm:"constraint:OnDelete:CASCADE;"`
	Carts          []_cartModel.Cart               `gorm:"constraint:OnDelete:CASCADE;"`
	Transactions   []_transactionModel.Transaction `gorm:"constraint:OnDelete:CASCADE;"`
}

func ModelToCore(data User) users.Core {
	return users.Core{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Username:       data.Username,
		Bio:            data.Bio,
		Role:           data.Role,
		Email:          data.Email,
		Password:       data.Password,
		Approvement:    data.Approvement,
	}
}

func CoreToModel(data users.Core) User {
	return User{
		Model:          gorm.Model{ID: data.ID},
		ProfilePicture: data.ProfilePicture,
		Username:       data.Username,
		Bio:            data.Bio,
		Role:           data.Role,
		Email:          data.Email,
		Password:       data.Password,
		Approvement:    data.Approvement,
	}
}
