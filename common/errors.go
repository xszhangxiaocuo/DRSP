package common

import "fmt"

type BError struct {
	Code BusinessCode
	Msg  string
}

func (e *BError) Error() string {
	return fmt.Sprintf("code:%v, msg:%s", e.Code, e.Msg)
}

func NewError(code BusinessCode, msg string) *BError {
	return &BError{
		Code: code,
		Msg:  msg,
	}
}
