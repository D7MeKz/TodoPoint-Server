package model

type Member struct {
	Id       int    `json:"Id"`
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	UserPw   string `json:"password"`
}
