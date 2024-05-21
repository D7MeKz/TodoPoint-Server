package data

type Me struct {
	Username string `json:"username"`
	ImgUrl   string `json:"imgUrl"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}
