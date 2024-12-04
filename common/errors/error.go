package errors

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	statusCode int
	msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.statusCode, e.msg)
}

func (e *HTTPError) GetStatusCode() int {
	return e.statusCode
}

func (e *HTTPError) GetErrorMsg() string {
	return e.msg
}

func NewBadRequestError(traceback string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusBadRequest,
		msg:        fmt.Sprintf("Bad Request: %s", traceback),
	}
}

func NewUnauthorizedError(msg string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusUnauthorized,
		msg:        fmt.Sprintf("Unauthorized: %s", msg),
	}
}

func NewForbiddenError(msg string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusForbidden,
		msg:        fmt.Sprintf("Forbidden: %s", msg),
	}
}

func NewNotFoundError(traceback string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusNotFound,
		msg:        fmt.Sprintf("Not Found: %s", traceback),
	}
}

func NewInvalidSlashCommandError(traceback string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusInternalServerError,
		msg:        fmt.Sprintf("Invalid Slash Command: %s", traceback),
	}
}

func NewInternalServerError(traceback string) *HTTPError {
	return &HTTPError{
		statusCode: http.StatusInternalServerError,
		msg:        fmt.Sprintf("Internal Server Error: %s", traceback),
	}
}
