package handler

import "alta-cookit-be/features/followers"

type FollowingRequest struct {
	ToUserID uint `json:"to_user_id" form:"to_user_id"`
}

func ReqToCore(data interface{}) *followers.FollowCore {
	res := followers.FollowCore{}

	switch data.(type) {
	case FollowingRequest:
		cnv := data.(FollowingRequest)
		res.ToUserID = cnv.ToUserID

	default:
		return nil
	}
	return &res
}
