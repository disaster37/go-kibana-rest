package kbapi

import (
	"fmt"
)

type APIError struct {
	Code    int
	Message string
}

func (e APIError) Error() string {
	return e.Message
}

func NewAPIError(code int, message string, params ...interface{}) APIError {
	return APIError{
		Code:    code,
		Message: fmt.Sprintf(message, params...),
	}
}
