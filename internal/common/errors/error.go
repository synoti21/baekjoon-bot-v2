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

func NewBadRequestError(msg ...string) *HTTPError {
	_msg := "Bad Request"
	if len(msg) > 0 {
		_msg = msg[0]
	}

	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Msg:        _msg,
	}
}

func NewUnauthorizedError(msg ...string) *HTTPError {
	_msg := "Unauthorized"
	if len(msg) > 0 {
		_msg = msg[0]
	}

	return &HTTPError{
		StatusCode: http.StatusUnauthorized,
		Msg:        _msg,
	}
}

func NewForbiddenError(msg ...string) *HTTPError {
	_msg := "Forbidden"
	if len(msg) > 0 {
		_msg = msg[0]
	}

	return &HTTPError{
		StatusCode: http.StatusForbidden,
		Msg:        _msg,
	}
}

func NewNotFoundError(msg ...string) *HTTPError {
	_msg := "Not Found"
	if len(msg) > 0 {
		_msg = msg[0]
	}
	return &HTTPError{
		StatusCode: http.StatusNotFound,
		Msg:        _msg,
	}
}

func NewInternalServerError(msg ...string) *HTTPError {
	_msg := "Internal Server Error"
	if len(msg) > 0 {
		_msg = msg[0]
	}
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Msg:        _msg,
	}
}
