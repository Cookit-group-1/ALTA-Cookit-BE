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
	showFollower := []Follower{}
	err := fq.db.Raw("SELECT f.id, f.from_user_id, f.to_user_id, u.username, u.profile_picture, u.role FROM followers f JOIN users u ON f.from_user_id = u.id where to_user_id = ?", userID).Find(&showFollower).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []followers.FollowCore{}, errors.New("no followers data found")
	}

	res := []followers.FollowCore{}
	for i := 0; i < len(showFollower); i++ {
		res = append(res, DataToCore(showFollower[i]))

	}
	if len(res) == 0 {
		return []followers.FollowCore{}, errors.New("no followers data found")
	}
	return res, nil
}

// ShowAllFollowing implements followers.FollowData
func (fq *FollowQuery) ShowAllFollowing(userID uint) ([]followers.FollowCore, error) {
	following := []Follower{}
	err := fq.db.Raw("SELECT f.id, f.from_user_id, f.to_user_id, u.username, u.profile_picture, u.role FROM followers f JOIN users u ON f.to_user_id = u.id where from_user_id= ?", userID).Scan(&following).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []followers.FollowCore{}, errors.New("no following data found")
	}

	res := []followers.FollowCore{}
	for i := 0; i < len(following); i++ {
		res = append(res, DataToCore(following[i]))

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
