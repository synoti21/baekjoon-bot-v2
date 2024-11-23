package slack

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/common"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/errors"
)

type slashCommandHandler struct {
	bot           bots.Interface
	signingSecret string
}

func NewRequestHandler() common.SlashCommandHandler {
	sb := NewSlackBot()
	secret := os.Getenv("SLACK_SIGNING_SECRET")

	return &slashCommandHandler{
		bot:           sb,
		signingSecret: secret,
	}
}

func (h *slashCommandHandler) HTTPServe() {
	port := os.Getenv("SLACK_PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(errors.ErrorHandlingMiddleware())
	r.POST("/receive", h.HandleSlashCommandReq)
	r.Run(":" + port)
}

func (h *slashCommandHandler) HandleSlashCommandReq(r *http.Request) {
	err := h.VerifySlashCommandReq(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	s, err := slack.SlashCommandParse(ctx.Request)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = common.RunSlashCommand(h.bot, common.SlashCommandReq{
		UserID:   s.UserID,
		Command:  s.Command,
		Argument: s.Text,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.String(http.StatusOK, fmt.Sprintf("Command %s executed successfully", s.Command))
}

func (h *slashCommandHandler) VerifySlashCommandReq(r *http.Request) error {
	v, err := slack.NewSecretsVerifier(r.Header, h.signingSecret)
	if err != nil {
		return errors.NewBadRequestError("Invalid request header")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.NewInternalServerError("Could not read request body")
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	_, err = v.Write(body)
	if err != nil {
		return errors.NewUnauthorizedError("Invalid Slack signature")
	}

	if err = v.Ensure(); err != nil {
		return errors.NewUnauthorizedError("Invalid Slack signature")
	}
	return nil
}
