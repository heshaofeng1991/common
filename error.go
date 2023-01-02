/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    error
	@Date    2022/4/19 11:07
	@Desc
*/

package internal

import (
	"errors"
	"fmt"
)

type Error struct {
	Msg  string
	Code int32
}

func (e Error) Error() string {
	return e.Msg
}

func (e Error) ErrorCode() int32 {
	return e.Code
}

// NewError .
func NewError(msg string, code int32) *Error {
	return &Error{
		Msg:  msg,
		Code: code,
	}
}

var ErrStack = errors.New("stack error")

func Stack(format string, args ...interface{}) error {
	exception := fmt.Sprintf(format, args) //nolint:govet

	return fmt.Errorf("%w - %s", ErrStack, exception)
}
