package discord

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/common"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/errors"
)

type slashCommandHandler struct {
	bot bots.Interface
}

func NewRequestHandler() (common.SlashCommandHandler, error) {
	sb := NewDiscordBot()
	return &slashCommandHandler{
		bot: sb,
	}, nil
}

func (h *slashCommandHandler) HTTPServe() {
	r := gin.Default()
	r.Use(errors.ErrorHandlingMiddleware())
	r.GET("/interactions", h.VerifySlashCommandReq)
	r.Run()
}

func (h *slashCommandHandler) VerifySlashCommandReq(ctx *gin.Context) {
	s, err := slack.SlashCommandParse(ctx.Request)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")
		return
	}

	if !s.ValidateToken() {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = h.HandleCommandReq(ctx, common.SlashCommandReq{
		Command:  s.Command,
		ChanID:   s.ChannelID,
		UserID:   s.UserID,
		Category: s.Text,
	})
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Not allowed slash commands")
		return
	}
}

func (h *slashCommandHandler) HandleCommandReq(ctx *gin.Context, scReq common.SlashCommandReq) error {
	switch scReq.Command {
	case "/prob":
		h.bot.SendProbToUser(scReq.UserID)
	case "/register":
		h.bot.RegisterUser(scReq.UserID)
	case "/category":
		categoryNum, err := strconv.Atoi(scReq.Category)
		if err != nil {
			return err
		}
		h.bot.SendProbToUserByCategory(scReq.UserID, consts.ProbCategory(categoryNum))
	}
	return nil
}
