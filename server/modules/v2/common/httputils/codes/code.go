package codes

type ErrorCode uint64

const (
	// 1xx is Bad Request error code
	InvalidHeader ErrorCode = 101
	InvalidBody   ErrorCode = 102
	InvalidQuery  ErrorCode = 103
	NotFound      ErrorCode = 104
	InvalidFile   ErrorCode = 105

	// 2xx is Unauthorized error code
	Unauthorized          ErrorCode = 200
	BadAuthenticationData ErrorCode = 201
	TokenExpired          ErrorCode = 202
	TokenCreateFailed     ErrorCode = 203

	// 3xx is Conflict error code
	AlreadyExist ErrorCode = 300

	// 3xx is Internal Server Error error code
	CreateFailed        ErrorCode = 401
	UpdateFailed        ErrorCode = 402
	DeleteFailed        ErrorCode = 403
	FindFail            ErrorCode = 404
	FileSaveFailed      ErrorCode = 405
	MetaDataUpdateError ErrorCode = 406
)

var errorMessage = map[ErrorCode]string{
	InvalidHeader: "The provided header values are invalid.",
	InvalidBody:   "The body of the request is invalid.",
	InvalidQuery:  "The query parameters are invalid.",
	NotFound:      "The requested resource was not found.",
	InvalidFile:   "The file is invalid.",

	BadAuthenticationData: "Authentication failed due to invalid credentials.",
	TokenExpired:          "The authentication token has expired.",
	TokenCreateFailed:     "The creation of the authentication token failed.",

	AlreadyExist: "The requested resource already exists.",

	CreateFailed:        "The creation of the requested resource failed.",
	UpdateFailed:        "The update of the requested resource failed.",
	DeleteFailed:        "The deletion of the requested resource failed.",
	FindFail:            "The requested resource was not found.",
	FileSaveFailed:      "The file save failed.",
	MetaDataUpdateError: "The metadata update failed.",
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
		return 409
	case 4:
		return 500
	default:
		return 500
	}
}
