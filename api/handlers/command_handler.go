package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/api/middlewares"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters/discord"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
)

type SlashCommandHandler struct {
	config  *config.HandlerConfig
	bot     bots.Interface
	adapter adapters.Interface
}

func New(cfg *config.HandlerConfig, b bots.Interface) (*SlashCommandHandler, error) {
	var adapter adapters.Interface

	switch cfg.Platform {
	case config.Slack:
		adapter = &slack.Adapter{}
	case config.Discord:
		adapter = &discord.Adapter{}
	default:
		return nil, errors.NewInternalServerError("Invalid adapter type")
	}

	return &SlashCommandHandler{
		bot:     b,
		config:  cfg,
		adapter: adapter,
	}, nil
}

func (h *SlashCommandHandler) Run() {
	port := os.Getenv("SLACK_PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(middlewares.VerifyRequestMiddleware(h.adapter, h.config.Secret))
	r.Use(middlewares.ErrorHandlingMiddleware())
	r.POST(h.config.RouteEndpoint(), h.SlashCommandHandlerFunc())
	r.Run(":" + port)
}

func (h *SlashCommandHandler) SlashCommandHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cmd, err := h.adapter.ParseSlashCommand(ctx.Request)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		resp, err := RunBotSlashCommand(h.bot, cmd)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		h.adapter.SendResponse(ctx, resp)
	}
}
