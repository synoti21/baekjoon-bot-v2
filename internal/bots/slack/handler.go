package slack

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots/common"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
)

type slashCommandHandler struct {
	bot           bots.Interface
	signingSecret string
}

func NewRequestHandler() common.SlashCommandHandler {
	sb := NewBot()
	return &slashCommandHandler{
		bot: sb,
	}
}

func (h *slashCommandHandler) HTTPServe() {
	r := gin.Default()
	r.POST("/receive", h.VerifyHTTPRequest)
	r.Run()
}

func (h *slashCommandHandler) VerifyHTTPRequest(ctx *gin.Context) {
	s, err := slack.SlashCommandParse(ctx.Request)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")
		return
	}

	scReq := common.SlashCommandReq{
		Command:  s.Command,
		ChanID:   s.ChannelID,
		UID:      s.UserID,
		Category: s.Text,
	}
	err = verifySignature(ctx, h.signingSecret)
	if err != nil {
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

func verifySignature(ctx *gin.Context, sign string) error {
	verifier, err := slack.NewSecretsVerifier(ctx.Request.Header, sign)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Could not create secrets verifier")
		return err
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not read request body")
		return err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	_, err = verifier.Write(body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not write body to verifier")
		return err
	}

	if err = verifier.Ensure(); err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return err
	}
	return nil
}
