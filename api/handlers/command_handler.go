package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/synoti21/baekjoon-slack-bot/api/middlewares"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters/discord"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
)

// SlashCommandHandler handles the slash command HTTP request, and executes the corresponding bot command.
type SlashCommandHandler struct {
	config  *config.HandlerConfig
	bot     bots.Interface
	adapter adapters.Interface
}

func New(cfg *config.HandlerConfig, b bots.Interface) (*SlashCommandHandler, error) {
	var adapter adapters.Interface

	// Select the adapter depending on the what platforms we are hosting.
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

func (h *SlashCommandHandler) Run() error {
	port := h.config.Port
	if _, err := strconv.ParseInt(port, 0, 64); err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Invalid config port: %v", port))
	}
	r := gin.Default()
	r.GET("/healthz", h.healthz())

	protected := r.Group("/")
	protected.Use(middlewares.VerifyRequestMiddleware(h.adapter, h.config.Secret))
	protected.Use(middlewares.ErrorHandlingMiddleware())

	protected.POST(h.config.RouteEndpoint(), h.SlashCommandHandlerFunc())

	return r.Run(":" + port)
}

// SlashCommandHandlerFunc is a handler function used by gin, to activate the slash command handler logics
func (h *SlashCommandHandler) SlashCommandHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cmd, err := h.adapter.ParseSlashCommand(ctx.Request)
		if err != nil {
			ctx.AbortWithError(err.GetStatusCode(), err)
			return
		}

		resp, err := h.RunBotSlashCommand(cmd)
		if err != nil {
			ctx.AbortWithError(err.GetStatusCode(), err)
			return
		}
		// Return the response to the user
		ctx.JSON(http.StatusOK, resp)
	}
}
