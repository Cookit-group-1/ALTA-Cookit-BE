package handler

import "alta-cookit-be/features/followers"

type UserResp struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Role           string `json:"role"`
}

type ListFollowingResponse struct {
	ID uint `json:"id"`
	// UserID   uint     `json:"user_id"`
	FromUserID     uint   `json:"from_user_id"`
	ToUserID       uint   `json:"to_user_id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Role           string `json:"role"`
}

func ToListFollowingResponse(data followers.FollowCore) ListFollowingResponse {
	return ListFollowingResponse{
		ID:             data.ID,
		FromUserID:     data.FromUserID,
		ToUserID:       data.ToUserID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
	}
}

type ListFollowerResponse struct {
	ID uint `json:"id"`
	// UserID         uint   `json:"user_id"`
	FromUserID     uint   `json:"from_user_id"`
	ToUserID       uint   `json:"to_user_id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Role           string `json:"role"`
}

func ToListFollowerResponse(data followers.FollowCore) ListFollowerResponse {
	return ListFollowerResponse{
		ID:             data.ID,
		FromUserID:     data.FromUserID,
		ToUserID:       data.ToUserID,
		Username:       data.Username,
		ProfilePicture: data.ProfilePicture,
		Role:           data.Role,
	}
}
