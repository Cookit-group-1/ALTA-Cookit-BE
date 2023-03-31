package handler

import "alta-cookit-be/features/followers"

type UserResp struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Role           string `json:"role"`
}

type ListFollowingResponse struct {
	ID       uint     `json:"id"`
	// UserID   uint     `json:"user_id"`
	ToUserID uint     `json:"to_user_id"`
	User     UserResp `json:"user"`
}

func ToListFollowingResponse(data followers.FollowCore) ListFollowingResponse {
	return ListFollowingResponse{
		ID:       data.UserID,
		// UserID:   data.UserID,
		ToUserID: data.ToUserID,
		User:     UserResp{UserID: data.UserID, Username: data.Username, ProfilePicture: data.ProfilePicture, Role: data.Role},
	}
}

type ListFollowerResponse struct {
	ID         uint     `json:"id"`
	UserID     uint     `json:"user_id"`
	FromUserID uint     `json:"from_user_id"`
	User       UserResp `json:"user"`
}

func ToListFollowerResponse(data followers.FollowCore) ListFollowerResponse {
	return ListFollowerResponse{
		ID:         data.UserID,
		UserID:     data.UserID,
		FromUserID: data.FromUserID,
		User:       UserResp{UserID: data.UserID, Username: data.Username, ProfilePicture: data.ProfilePicture, Role: data.Role},
	}
}
