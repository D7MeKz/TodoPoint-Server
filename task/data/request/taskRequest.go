package request

type CreateTask struct {
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
}
