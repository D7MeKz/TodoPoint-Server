package webutils

import "net/http"

/*
일반적인 오류 : 1000 ~ 1999
Member : 2000 ~ 2999
Task : 3000 ~ 3999
*100 -> DB 오류
*/
// Error Code
const (
	SUCCESS       = 0
	ERROR_GO_CODE = 1001
	INVALID_JSON  = 1002
	// Member
	INVALID_MEMBER  = 2001
	ERROR_MEMBER_DB = 2101

	// Task
	ERROR_TASK_DB = 3101
)

type ErrResponse struct {
	Message    string
	StatusCode int
}

var errMsg = map[int]ErrResponse{
	ERROR_GO_CODE:   {Message: "Logic Error", StatusCode: http.StatusBadRequest},
	INVALID_JSON:    {Message: "Invalid Json", StatusCode: http.StatusBadRequest},
	INVALID_MEMBER:  {Message: "Invalid Member", StatusCode: http.StatusUnauthorized},
	ERROR_MEMBER_DB: {Message: "Member DB Error", StatusCode: http.StatusInternalServerError},
	ERROR_TASK_DB:   {Message: "Task DB Error", StatusCode: http.StatusInternalServerError},
}

func getErrorMsg(code int) ErrResponse {
	msg, ok := errMsg[code]
	if ok {
		return msg
	}
	return errMsg[code]
}
