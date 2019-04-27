package handler

import (
	"fmt"
)

type ApplicationError struct {
	Code    int
	Message string
}

func NewApplicationError(code int, message string) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Message: message,
	}
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("HttpStatus: %d, Message: %s", e.Code, e.Message)
}
