package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type SlashCommandRequest struct {
	UserID    string
	ChannelID string
	Command   consts.SlashCommand
	Arg       string
}

type Interface interface {
	VerifyRequest(r *http.Request, signature string) error
	ParseSlashCommand(r *http.Request) (*SlashCommandRequest, error)
	CreateMessage(prob *schema.BaekjoonProb) (interface{}, error)
	SendResponse(ctx *gin.Context, response interface{}) error
}
