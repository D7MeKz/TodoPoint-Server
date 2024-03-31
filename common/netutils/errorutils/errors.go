package errorutils

import (
	"errors"
	"net/http"
)

type ErrorCode uint

type NetError struct {
	Code ErrorCode
	Err  error
}

func NewNetError(code ErrorCode, err error) *NetError {
	return &NetError{Code: code, Err: err}
}

func IsNetError(err error) bool {
	if err == nil {
		return false
	}
	var e *NetError
	return errors.As(err, &e)
}

func (e *NetError) GetStatus() ErrorCode {
	// 140400 (1 : label, 404: status code, 00 : Meta)
	deletedMeta := e.GetCode() / 100
	code := deletedMeta % 1000
	return code
}

func (e *NetError) GetCode() ErrorCode {
	return e.Code
}

func (e *NetError) Error() string {
	return e.Err.Error()
}

func (e *NetError) Unwrap() error {
	return e.Err
}

// error to NetError
func ConvertNetError(err error) (*NetError, bool) {
	for err != nil {
		switch {
		case IsNetError(err):
			return err.(*NetError), true
		}
		err = errors.Unwrap(err)
	}
	return nil, false
}
func IsBadRequest(err NetError) bool {
	status := err.GetStatus()
	if status != http.StatusBadRequest {
		return false
	}
	return true
}

func IsNotFound(err NetError) bool {
	status := err.GetStatus()
	if status != http.StatusNotFound {
		return false
	}
	return true
}

func IsInternalServerError(err NetError) bool {
	status := err.GetStatus()
	if status != http.StatusInternalServerError {
		return false
	}
	return true
}
