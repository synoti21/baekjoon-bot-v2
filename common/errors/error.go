package errors

import (
	"fmt"
	"net/http"
)

type BaseError struct {
	statusCode int
	msg        string
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%d: %s", e.statusCode, e.msg)
}

func (e *BaseError) GetStatusCode() int {
	return e.statusCode
}

func (e *BaseError) GetErrorMsg() string {
	return e.msg
}

func NewBadRequestError(traceback string) *BaseError {
	return &BaseError{
		statusCode: http.StatusBadRequest,
		msg:        fmt.Sprintf("Bad Request: %s", traceback),
	}
}

func NewUnauthorizedError(msg string) *BaseError {
	return &BaseError{
		statusCode: http.StatusUnauthorized,
		msg:        fmt.Sprintf("Unauthorized: %s", msg),
	}
}

func NewForbiddenError(msg string) *BaseError {
	return &BaseError{
		statusCode: http.StatusForbidden,
		msg:        fmt.Sprintf("Forbidden: %s", msg),
	}
}

func NewNotFoundError(traceback string) *BaseError {
	return &BaseError{
		statusCode: http.StatusNotFound,
		msg:        fmt.Sprintf("Not Found: %s", traceback),
	}
}

func NewInvalidSlashCommandError(traceback string) *BaseError {
	return &BaseError{
		statusCode: http.StatusInternalServerError,
		msg:        fmt.Sprintf("Invalid Slash Command: %s", traceback),
	}
}

func NewInternalServerError(traceback string) *BaseError {
	return &BaseError{
		statusCode: http.StatusInternalServerError,
		msg:        fmt.Sprintf("Internal Server Error: %s", traceback),
	}
}
