package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/errors"
)

type SlashCommandReq struct {
	UserID    string
	ChannelID string
	Command   string
	Argument  string
}

type SlashCommandHandler interface {
	HTTPServe()
	HandleSlashCommandReq(r *http.Request) error
	VerifySlashCommandReq(r *http.Request) error
}

func RunBotSlashCommand(b bots.Interface, r SlashCommandReq) error {
	switch r.Command {
	case "/register":
		return b.RegisterUser(r.UserID, r.ChannelID)
	case "/prob":
		return b.SendProbToUser(r.UserID)
	case "/category":
		categoryNum, err := strconv.Atoi(r.Argument)
		if err != nil {
			return err
		}
		if err := consts.ValidateProbCategory(categoryNum); err != nil {
			return err
		}
		return b.SendProbToUserByCategory(r.UserID, r.ChannelID, consts.ProbCategory(categoryNum))
	case "/daily":
		t, err := time.Parse(r.Argument, "15 04")
		if err != nil {
			return errors.NewInvalidSlashCommandError("Invalid time format")
		}
		return b.ScheduleDailyProb(r.UserID, r.ChannelID, t)
	case "/categorylist":
		return b.ShowProbCategoryList(r.UserID, r.ChannelID)
	case "/help":
		return b.ShowHelpGuide(r.UserID, r.ChannelID)
	default:
		return errors.NewInvalidSlashCommandError("Unknown slash command")
	}
}
