package data

type CreateReq struct {
	Title string `json:"title"`
}

type AddSubReq struct {
	TaskId string `json:"task_id"`
	Title  string `json:"title"`
}
