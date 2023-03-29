package data

import (
	"alta-cookit-be/features/followers"
	"errors"
	"log"

	"gorm.io/gorm"
)

func New(db *gorm.DB) followers.FollowData {
	return &FollowQuery{
		db: db,
	}
}

type FollowQuery struct {
	db *gorm.DB
}

// Follow implements followers.FollowData
func (fq *FollowQuery) Follow(userID, followingID uint) error {
	following := Follower{
		FromUserID: userID,
		ToUserID: followingID,
	}
	followQry := fq.db.Create(&following)
	rowAffect := followQry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("your already followed this user")
	}

	err := followQry.Error
	if err != nil {
		log.Println("follow query error", err.Error())
		return errors.New("follow user has been fail")
	}

	return nil
}

// ShowAllFollower implements followers.FollowData
func (fq *FollowQuery) ShowAllFollower() ([]followers.FollowCore, error) {
	panic("unimplemented")
}

// ShowAllFollowing implements followers.FollowData
func (fq *FollowQuery) ShowAllFollowing() ([]followers.FollowCore, error) {
	panic("unimplemented")
}

// Unfollow implements followers.FollowData
func (fq *FollowQuery) Unfollow(userID, followingID uint) error {
	res := Follower{
		FromUserID: userID,
		ToUserID: followingID,
	}
	unfollowQry := fq.db.Delete(&res)

	rowAffect := unfollowQry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user to unfollow")
	}

	err := unfollowQry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("unfollow account has been fail")
	}

	return nil
}
