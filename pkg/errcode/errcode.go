package errcode

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	Code    int
	Msg     string
	Details []string
	Err     error
}

var codes = map[int]string{}

func NewError(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换", code))
	}
	codes[code] = msg
	return Error{Code: code, Msg: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Msg)
}

func (e Error) ErrorWithDetails() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s, 详细信息：%s", e.Code, e.Msg, strings.Join(e.Details, ";"))
}

func (e Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e Error) WithDetails(details ...string) Error {
	e.Details = append(e.Details, details...)
	return e
}

func (e Error) WithError(err error) Error {
	e.Err = err
	return e
}

func IsError(err error) bool {
	_, ok := err.(Error)
	for !ok {
		in := errors.Unwrap(err)
		if in == nil {
			return false
		}
		_, ok = err.(Error)
	}
	return true
}

func (e Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InValidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code, UnauthorizedTokenError.Code, UnauthorizedTokenTimeOut.Code:
		return http.StatusUnauthorized
	case TooManyRequest.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
