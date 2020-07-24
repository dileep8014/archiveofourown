package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Msg     string
	Details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Msg)
}

func (e *Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.Details = append(e.Details, details...)
	return e
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InValidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code, UnauthorizedTokenError.Code, UnauthorizedTokenGenerate.Code, UnauthorizedTokenTimeOut.Code:
		return http.StatusUnauthorized
	case TooManyRequest.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
