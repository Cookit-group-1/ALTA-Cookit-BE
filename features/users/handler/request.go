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
	default:
		return nil
	}
	return &res
}
