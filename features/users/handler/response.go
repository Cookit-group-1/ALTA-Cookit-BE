package handler

import "alta-cookit-be/features/users"

type ProfileResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
	Bio            string `json:"bio"`
}

func ToProfileResponse(data users.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Username:       data.Username,
		Bio:            data.Bio,
	}
}
