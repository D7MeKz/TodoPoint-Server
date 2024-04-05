package codes

type WebCode uint

const (
	// Common
	GlobalInternalServerError WebCode = 50000

	// Member
	// 200 : OK
	MemberOK           WebCode = 120000
	MemberLoginSuccess WebCode = 120001
	// 201 : Created
	MemberCreationSuccess WebCode = 120101
	// 400 : Bad Request
	MemberInvalidJson WebCode = 140001
	MemberInvalidUri  WebCode = 140002
	// 404 : Not Found
	MemberNotFound WebCode = 140400

	MemberInternalServerError         = 150000
	MemberCreationError       WebCode = 150001

	TaskInvaliJson WebCode = 240001

	TaskCreationError WebCode = 250001
)

func GetStatus(c WebCode) int {
	// 140400 (1 : label, 404: status codes, 00 : Meta)
	deletedMeta := c / 100
	code := deletedMeta % 1000
	return int(code)
}
