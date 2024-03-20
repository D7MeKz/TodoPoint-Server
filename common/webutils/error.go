package webutils

import "net/http"

// Error Code
const (
	INVALID_MEMBER = 200000
)

type errRes struct {
	message    string
	statusCode int
}

var msg = map[int]errRes{
	INVALID_MEMBER: errRes{message: "Invalid Member", statusCode: http.StatusUnauthorized},
}

func GetErrorMsg(code int) (string, int) {
	msg, ok := msg[code]
	if ok {
		return
	}
}
