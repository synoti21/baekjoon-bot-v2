package common

import (
	"net/http"
	"strconv"

	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/errors"
)

type SlashCommandReq struct {
	UserID   string
	Command  string
	Argument string
}

type SlashCommandHandler interface {
	HandleSlashCommandReq(req *http.Request)
	VerifySlashCommandReq(req *http.Request) error
}

func RunSlashCommand(b bots.Interface, r SlashCommandReq) error {
	switch r.Command {
	case "/prob":
		return b.SendProbToUser(r.UserID)
	case "/register":
		return b.RegisterUser(r.UserID)
	case "/category":
		categoryNum, err := strconv.Atoi(r.Argument)
		if err != nil {
			return err
		}
		return b.SendProbToUserByCategory(r.UserID, consts.ProbCategory(categoryNum))
	default:
		return errors.NewBadRequestError("Unknown Slash Command")
	}
}
