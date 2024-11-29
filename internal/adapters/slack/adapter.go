package slack

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Adapter struct{}

var _ adapters.Interface = (*Adapter)(nil)

func (a *Adapter) VerifyRequest(r *http.Request, secret string) error {
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

func (a *Adapter) ParseSlashCommand(r *http.Request) (*adapters.SlashCommandRequest, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) CreateMessage(prob *schema.BaekjoonProb) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) SendResponse(ctx *gin.Context, response interface{}) error {
	panic("not implemented") // TODO: Implement
}
