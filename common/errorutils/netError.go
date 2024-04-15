package errorutils

import (
	"todopoint/common/errorutils/codes"
)

type NetError struct {
	Code codes.WebCode
	Err  error
}

func NewNetError(code codes.WebCode, err error) error {
	return &NetError{Code: code, Err: err}
}

func (e *NetError) GetCode() codes.WebCode {
	return e.Code
}
func (e *NetError) Error() string {
	return e.Err.Error()
}
