package errors

import (
	"fmt"
)

type Error struct {
	Code    int `json:"code"`
	Message int `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error (%d): %s", e.Code, e.Message)
}

func newError(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

const (
	UnknownErr ErrorCode = iota + 90001
)

const (
	InvalidRequestParams ErrorCode = iota + 10001
	MissingRequestParams
	UnsupportedLoginType
)

const (
	UserNotFound ErrorCode = iota + 2000
	InvalidPassword
)

var (
	// common
	ErrUnknown = newError(UnknownErr, "unknown error")
	// requeset
	ErrInvalidRequestParams = newError(InvalidRequestParams, "invalid request params")
	ErrMissingRequestParams = newError(MissingRequestParams, "missing request params")
	ErrUnsupportedLoginType = newError(UnsupportedLoginType, "unsupported login type")
	ErrUserNotFound         = newError(UserNotFound, "user not found")
	ErrInvalidPassword      = newError(InvalidPassword, "invalid password")
)
