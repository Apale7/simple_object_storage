package model

import user_center "Apale7/simple_object_storage/proto/user-center"

type LoginReq struct {
	Base
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterReq struct {
	Base
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Nickname    string `json:"nickname" form:"nickname"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type LoginResp struct {
	user_center.LoginResponse
	Auth []string `json:"auth"`
}
