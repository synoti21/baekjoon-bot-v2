package discord

import (
	"net/http"

	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/common"
)

type slashCommandHandler struct {
	bot bots.Interface
}

func NewRequestHandler() (common.SlashCommandHandler, error) {
	db := NewDiscordBot()
	return &slashCommandHandler{
		bot: db,
	}, nil
}

func (h *slashCommandHandler) HTTPServe() {
	panic("not implemented") // TODO: Implement
}

func (h *slashCommandHandler) HandleSlashCommandReq(r *http.Request) error {
	panic("not implemented") // TODO: Implement
}

func (h *slashCommandHandler) VerifySlashCommandReq(r *http.Request) error {
	panic("not implemented") // TODO: Implement
}
