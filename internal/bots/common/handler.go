package common

import (
	"github.com/gin-gonic/gin"
)

type SlashCommandReq struct {
	Command  string
	ChanID   string
	UID      string
	Category string
}

type SlashCommandHandler interface {
	HTTPServe()
	VerifyHTTPRequest(ctx *gin.Context)
	HandleCommandRequest(ctx *gin.Context, scReq SlashCommandReq) error
}
