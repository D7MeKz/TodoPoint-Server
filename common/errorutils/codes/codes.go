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
	//NOTE : 200
	TaskOneSuccess    WebCode = 220000
	TaskListSuccess   WebCode = 220001
	TaskUpdateSuccess WebCode = 2200002

	SubtaskOneSuccess  WebCode = 220050
	SubtaskListSuccess WebCode = 220051

	// NOTE : 201
	TaskCreationSuccess    WebCode = 220101
	SubtaskCreationSuccess WebCode = 220151

	// NOTE : 400
	TaskInvalidJson  WebCode = 240001
	TaskInvalidUri   WebCode = 240002
	TaskInvalidQuery WebCode = 240003
	TaskDoesNotFound WebCode = 240400

	SubtaskInvalidJson  WebCode = 240051
	SubtaskDoesNotFound WebCode = 240050
	// NOTE : 500
	TaskCreationError WebCode = 250001
	TaskListError     WebCode = 250002
	TaskDecodingErr   WebCode = 250003

	SubtaskCreationErr WebCode = 250051
	SubtaskAdditionErr WebCode = 250054
	SubtaskUpdateErr   WebCode = 250055

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
