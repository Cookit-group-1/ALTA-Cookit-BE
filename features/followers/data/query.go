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
		ToUserID:   followingID,
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
func (fq *FollowQuery) ShowAllFollower(userID uint) ([]followers.FollowCore, error) {
	follower := []Follower{}
	log.Println(follower)
	err := fq.db.Raw("SELECT u.id, u.username, u.profile_picture, u.role, f.to_user_id FROM users u JOIN followers f ON u.id = f.from_user_id WHERE u.id = ?", userID).Scan(&follower).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []followers.FollowCore{}, errors.New("no following data found")
	}

	res := []followers.FollowCore{}
	for i := 0; i < len(follower); i++ {
		res = append(res, DataToCore(follower[i]))

	}
	if len(res) == 0 {
		return []followers.FollowCore{}, errors.New("no following data found")
	}
	return res, nil
}

// ShowAllFollowing implements followers.FollowData
func (fq *FollowQuery) ShowAllFollowing(userID uint) ([]followers.FollowCore, error) {
	following := []Followings{}
	log.Println(following)
	err := fq.db.Raw("SELECT u.id as user_id, u.username, u.profile_picture, u.role, f.to_user_id as following_user_id FROM users u LEFT JOIN followers f ON u.id = f.from_user_id WHERE u.id = ?", userID).Scan(&following).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []followers.FollowCore{}, errors.New("no following data found")
	}

	res := []followers.FollowCore{}
	for i := 0; i < len(following); i++ {
		res = append(res, FollowingDataToCore(following[i]))

	}
	if len(res) == 0 {
		return []followers.FollowCore{}, errors.New("no following data found")
	}
	return res, nil
}

// Unfollow implements followers.FollowData
func (fq *FollowQuery) Unfollow(userID, followingID uint) error {
	res := Follower{
		FromUserID: userID,
		ToUserID:   followingID,
	}
	unfollowQry := fq.db.Unscoped().Where("from_user_id = ? AND to_user_id = ?", userID, followingID).Delete(&res)

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
