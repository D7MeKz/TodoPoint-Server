package main

type Member struct {
	Id       int    `json:"Id"`
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	UserPw   string `json:"password"`
}

// Member list
var members map[int]Member

var lastId int
