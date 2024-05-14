package httputils

type ErrorCode uint64

const (
	// 1xx is Bad Request error code
	InvalidHeader ErrorCode = 101
	InvalidBody   ErrorCode = 102
	InvalidQuery  ErrorCode = 103
	NotFound      ErrorCode = 104

	// 2xx is Unauthorized error code
	BadAuthenticationData ErrorCode = 201
	TokenExpired          ErrorCode = 202
	TokenCreateFailed     ErrorCode = 203

	// 3xx is Internal Server Error error code
	CreateFailed ErrorCode = 301
	UpdateFailed ErrorCode = 302
	DeleteFailed ErrorCode = 303
)

var errorMessage = map[ErrorCode]string{
	InvalidHeader: "The provided header values are invalid.",
	InvalidBody:   "The body of the request is invalid.",
	InvalidQuery:  "The query parameters are invalid.",
	NotFound:      "The requested resource was not found.",

	BadAuthenticationData: "Authentication failed due to invalid credentials.",
	TokenExpired:          "The authentication token has expired.",
	TokenCreateFailed:     "The creation of the authentication token failed.",

	CreateFailed: "The creation of the requested resource failed.",
	UpdateFailed: "The update of the requested resource failed.",
	DeleteFailed: "The deletion of the requested resource failed.",
}

func GetErrorMsg(code ErrorCode) string {
	return errorMessage[code]
}

// ParseStatusCode returns the status code of the error
// If ErrorCode is 40001, it returns 400
func ParseStatusCode(code ErrorCode) int {
	flag := code / 100
	switch flag {
	case 1:
		return 400
	case 2:
		return 401
	case 3:
		return 500
	default:
		return 500
	}
}
