package errors

import "net/http"

func NewInvalidDiscordSlashCommmandError() *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Msg:        "Invalid Discord Slash Command",
	}
}
