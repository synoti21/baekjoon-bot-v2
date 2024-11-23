package slack

import (
	"bytes"
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
	r.POST("/receive", h.ginHandlerFunc)
	r.Run(":" + port)
}

func (h *slashCommandHandler) ginHandlerFunc(c *gin.Context) {
	err := h.HandleSlashCommandReq(c.Request)
	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, "Command executed successfully")
}

func (h *slashCommandHandler) HandleSlashCommandReq(r *http.Request) error {
	err := h.VerifySlashCommandReq(r)
	if err != nil {
		return err
	}

	s, err := slack.SlashCommandParse(r)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = common.RunBotSlashCommand(h.bot, common.SlashCommandReq{
		UserID:   s.UserID,
		Command:  s.Command,
		Argument: s.Text,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *slashCommandHandler) VerifySlashCommandReq(r *http.Request) error {
	v, err := slack.NewSecretsVerifier(r.Header, h.signingSecret)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	_, err = v.Write(body)
	if err != nil {
		return errors.NewUnauthorizedError(err.Error())
	}

	if err = v.Ensure(); err != nil {
		return errors.NewUnauthorizedError(err.Error())
	}
	return nil
}
