package errors

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Msg)
}

func NewBadRequestError(traceback string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Msg:        fmt.Sprintf("Bad Request: %s", traceback),
	}
}

func NewUnauthorizedError(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusUnauthorized,
		Msg:        fmt.Sprintf("Unauthorized: %s", msg),
	}
}

func NewForbiddenError(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusForbidden,
		Msg:        fmt.Sprintf("Forbidden: %s", msg),
	}
}

func NewNotFoundError(traceback string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusNotFound,
		Msg:        fmt.Sprintf("Not Found: %s", traceback),
	}
}

func NewInvalidSlashCommandError(traceback string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Msg:        fmt.Sprintf("Invalid Slash Command: %s", traceback),
	}
}

func NewInternalServerError(traceback string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Msg:        fmt.Sprintf("Internal Server Error: %s", traceback),
	}
}
