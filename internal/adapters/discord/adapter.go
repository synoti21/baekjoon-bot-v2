package discord

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Adapter struct{}

var _ adapters.Interface = (*Adapter)(nil)

func (a *Adapter) VerifyRequest(r *http.Request, secret string) error {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) ParseSlashCommand(r *http.Request) (*adapters.SlashCommandRequest, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) CreateMessage(prob *schema.BaekjoonProb) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) SendResponse(ctx *gin.Context, response interface{}) error {
	panic("not implemented") // TODO: Implement
}
