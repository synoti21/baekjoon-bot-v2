package discord

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/common"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
)

type slashCommandHandler struct {
	bot bots.Interface
}

func NewRequestHandler() (common.SlashCommandHandler, error) {
	sb := NewBot()
	return &slashCommandHandler{
		bot: sb,
	}, nil
}

func (h *slashCommandHandler) HTTPServe() {
	r := gin.Default()
	r.GET("/interactions", h.VerifyHTTPRequest)
	r.Run()
}

func (h *slashCommandHandler) VerifyHTTPRequest(ctx *gin.Context) {
	s, err := slack.SlashCommandParse(ctx.Request)
	scReq := common.SlashCommandReq{
		Command:  s.Command,
		ChanID:   s.ChannelID,
		UID:      s.UserID,
		Category: s.Text,
	}

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")
		return
	}
	if !s.ValidateToken() {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	err = h.HandleCommandRequest(ctx, scReq)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Not allowed slash commands")
		return
	}
}

func (h *slashCommandHandler) HandleCommandRequest(ctx *gin.Context, scReq common.SlashCommandReq) error {
	switch scReq.Command {
	case "/prob":
		h.bot.SendProbToUser(scReq.UID)
	case "/register":
		h.bot.RegisterUser(scReq.UID)
	case "/category":
		categoryNum, err := strconv.Atoi(scReq.Category)
		if err != nil {
			return err
		}
		h.bot.SendProbToUserByCategory(scReq.UID, consts.ProbCategory(categoryNum))
	}
	panic("not implemented") // TODO: Implement
}
