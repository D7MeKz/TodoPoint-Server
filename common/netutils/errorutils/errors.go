package errorutils

import (
	"errors"
	"net/http"
	"todopoint/common/netutils/codes"
)

type NetError struct {
	Code codes.WebCode
	Err  error
}

func NewNetError(code codes.WebCode, err error) error {
	return &NetError{Code: code, Err: err}
}

func IsNetError(err error) bool {
	if err == nil {
		return false
	}
	var e *NetError
	return errors.As(err, &e)
}

func Convert(err error) (*NetError, bool) {
	for err != nil {
		var netError *NetError
		switch {
		case errors.As(err, &netError):
			return netError, true
		}
		//err = errors.Unwrap(err)
	}
	return nil, false
}

func (e *NetError) GetCode() codes.WebCode {
	return e.Code
}
func (e *NetError) Error() string {
	return e.Err.Error()
}

//	func (e *NetError) Unwrap() error {
//		return e.Err
//	}
//
// error to NetError

func IsBadRequest(err NetError) bool {
	status := codes.GetStatus(err.Code)
	if status != http.StatusBadRequest {
		return false
	}
	return true
}

func IsNotFound(err NetError) bool {
	status := codes.GetStatus(err.Code)
	if status != http.StatusNotFound {
		return false
	}
	return true
}

func IsInternalServerError(err NetError) bool {
	status := codes.GetStatus(err.Code)
	if status != http.StatusInternalServerError {
		return false
	}
	return true
}