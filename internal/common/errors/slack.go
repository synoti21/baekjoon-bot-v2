package errors

import "net/http"

func NewInvalidSlackSlashCommmandError() *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Msg:        "Invalid Slack Slash Command",
	}
}
