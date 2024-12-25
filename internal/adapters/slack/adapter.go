package slack

import (
	"bytes"
	"io"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Adapter struct{}

var _ adapters.Interface = (*Adapter)(nil)

func (a *Adapter) VerifyRequest(r *http.Request, secret string) *errors.BaseError {
	v, err := slack.NewSecretsVerifier(r.Header, secret)
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

func (a *Adapter) ParseSlashCommand(r *http.Request) (*adapters.SlashCommandRequest, *errors.BaseError) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	sc, ok := consts.ValidateSlashCommand(s.Command)
	if !ok {
		return nil, errors.NewBadRequestError("Invalid Slack command")
	}

	return &adapters.SlashCommandRequest{
		Command:   sc,
		ChannelID: s.ChannelID,
		UserID:    s.UserID,
		Arg:       s.Text,
	}, nil
}

func (a *Adapter) CreateTextMessage(text string) (interface{}, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) CreateCategoryListMessage() (interface{}, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) CreateHelpGuideMessage() (interface{}, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) CreateProblemMessage(prob *schema.BaekjoonProb) (interface{}, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
}
