package discord

import (
	"net/http"

	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Adapter struct{}

var _ adapters.Interface = (*Adapter)(nil)

func (a *Adapter) VerifyRequest(r *http.Request, secret string) *errors.BaseError {
	panic("not implemented") // TODO: Implement
}

func (a *Adapter) ParseSlashCommand(r *http.Request) (*adapters.SlashCommandRequest, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
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
