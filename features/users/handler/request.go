package handler

import "alta-cookit-be/features/users"

type RegisterReq struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateProfileReq struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Username       string `json:"username" form:"username"`
	Bio            string `json:"bio" form:"bio"`
}

func ReqToCore(data interface{}) *users.Core {
	res := users.Core{}

	switch data.(type) {
	case RegisterReq:
		cnv := data.(RegisterReq)
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	case LoginReq:
		cnv := data.(LoginReq)
		res.Username = cnv.Username
		res.Password = cnv.Password
	case UpdateProfileReq:
		cnv := data.(UpdateProfileReq)
		res.ProfilePicture = cnv.ProfilePicture
		res.Username = cnv.Username
		res.Bio = cnv.Bio
	default:
		return nil
	}
	return &res
}
