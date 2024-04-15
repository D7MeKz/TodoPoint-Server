package data

type TaskId struct {
	Id string `json:"task_id"`
}

type TaskInfo struct {
	TaskId    string `json:"task_id"`
	CreatedAt string `json:"created_at"`
	Status    bool   `json:"status"`
}
