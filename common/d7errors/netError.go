package d7errors

import (
	"github.com/sirupsen/logrus"
	"todopoint/common/d7errors/codes"
)

type NetError struct {
	Code codes.WebCode
	Err  error
}

func NewNetError(code codes.WebCode, err error) *NetError {
	logrus.Errorf("Code : %d, Error : %v", code, err)
	return &NetError{Code: code, Err: err}
}

func (e *NetError) GetCode() codes.WebCode {
	return e.Code
}
func (e *NetError) Error() string {
	return e.Err.Error()
}
