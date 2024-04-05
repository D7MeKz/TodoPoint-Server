package req

type CreateReq struct {
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
}
