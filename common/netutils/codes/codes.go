package consts

type WebCode uint

const (
	// Common
	GlobalInternalServerError WebCode = 50000

	// Member
	MemberCreationSuccess WebCode = 120101
	MemberIsNotFound      WebCode = 140400
	MemberInvalidJson     WebCode = 140401

	MemberCreationError WebCode = 150001
)
