package response

import (
	"todopoint/common/errorutils/codes"
)

// Response
type ErrorResponse struct {
	Code codes.WebCode `json:"codes"`
}

func NewErrorResponse(code codes.WebCode) *ErrorResponse {
	return &ErrorResponse{Code: code}
}
