package codes

type WebCode uint

const (
	ZeroCode WebCode = 0

	// Common
	GlobalInternalServerError WebCode = 50000

	// Member
	// NOTE : 200(OK)
	MemberOK           WebCode = 120000
	MemberLoginSuccess WebCode = 120001
	// 201 : Created
	MemberCreationSuccess WebCode = 120101
	// 400 : Bad Request
	MemberInvalidJson WebCode = 140001
	MemberInvalidUri  WebCode = 140002

	// 401
	TokenExpired     WebCode = 140100
	TokenCreationErr WebCode = 140101

	// 404 : Not Found
	MemberNotFound WebCode = 140400

	MemberInternalServerError WebCode = 150000
	MemberCreationError       WebCode = 150001
	TokenCreationError        WebCode = 150002
	TokenExpiredErr           WebCode = 150003
	// ---------- Task ---------------
	// NOTE : 200
	TaskCreationSuccess WebCode = 220101

	// NOTE : 400
	TaskInvalidJson WebCode = 240001

	// NOTE : 500
	TaskCreationError WebCode = 250001

	// NOTE : 503
	TaskMemberUnavailable WebCode = 250301
)

/*
GetStatus Function
NOTE : 140400 (1 : label, 404: status codes, 00 : Meta)
*/
func GetStatus(c WebCode) int {

	deletedMeta := c / 100
	code := deletedMeta % 1000
	return int(code)
}

// ConvertFrom
func ConvertFrom(a any) WebCode {
	switch v := a.(type) {
	case float64:
		code := WebCode(v)
		return code
	default:
		panic("Unexpected error to convert WebCode")
		return ZeroCode
	}
}
