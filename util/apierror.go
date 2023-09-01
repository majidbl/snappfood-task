package util

import (
	"fmt"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}
func (e ApiError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func CastError(err error) ApiError {
	e, ok := err.(ApiError)
	if ok {
		return e
	}

	return ApiError{
		Code:    500,
		Message: "internal error",
	}
}
