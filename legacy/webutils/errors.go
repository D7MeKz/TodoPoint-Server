package webutils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ErrorType int

type Error struct {
	// Code is a custom error codes
	ErrorType ErrorType
	// Err is a error string
	Err error
	// Description is a human-friendly message.
	Description string
}

// ErrorRes is a response object when error happens.
type ErrorRes struct {
	// Code is http status codes
	Code int `json:"codes"`
}

// ErrorType
/*
	일반적인 오류 : 1000 ~ 1999
	Member : 2000 ~ 2999
	Task : 3000 ~ 3999
	Point : 4000 ~ 4999

	-- Rule --
	*1** 도메인 정보와 관련된 오류
	*2** 번째는 DB와 관련된 오류

	***0 : 성공 시 부여

*/
const (
	// Common
	SUCCESS             ErrorType = 0
	INVALID_JSON_FORMAT ErrorType = 1001
	INVALID_URI_FORMAT  ErrorType = 1002

	// Member
	INVALID_MEMBER   ErrorType = 2001
	MEMBER_NOT_FOUND ErrorType = 2002
	LOGIN_FAILED     ErrorType = 2003
	ERROR_MEMBER_DB  ErrorType = 2101

	// Task
	ERROR_TASK_DB ErrorType = 3101
)

var codeMap = map[ErrorType]int{
	// Common
	INVALID_JSON_FORMAT: http.StatusBadRequest,
	INVALID_URI_FORMAT:  http.StatusBadRequest,

	// Member
	INVALID_MEMBER:   http.StatusUnauthorized,
	MEMBER_NOT_FOUND: http.StatusNotFound,
	LOGIN_FAILED:     http.StatusNotFound,
	ERROR_MEMBER_DB:  http.StatusInternalServerError,

	// Task
	ERROR_TASK_DB: http.StatusInternalServerError,
}

// getCode is get Status codes from codeMap.
func getCode(flag ErrorType) *ErrorRes {
	return &ErrorRes{Code: codeMap[flag]}
}

// Error return error message.
func (msg *Error) Error() string {
	return msg.Err.Error()
}

func NewError(errorType ErrorType, err error) *Error {
	return &Error{ErrorType: errorType, Err: err, Description: ""}
}

func ErrorFunc(ctx *gin.Context, err *Error) {
	res := getCode(err.ErrorType)
	log.Println(err)

	ctx.AbortWithStatusJSON(res.Code, res)
	return
}
