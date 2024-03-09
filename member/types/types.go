package types

type Member struct {
	Id       int    `json:"Id"`
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	UserPw   string `json:"password"`
}

var lastId int
