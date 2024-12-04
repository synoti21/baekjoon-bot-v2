package adapters

import (
	"net/http"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Interface interface {
	VerifyRequest(r *http.Request, signature string) *errors.HTTPError
	ParseSlashCommand(r *http.Request) (*SlashCommandRequest, *errors.HTTPError)
	CreateProblemMessage(prob *schema.BaekjoonProb) (interface{}, *errors.HTTPError)
	CreateCategoryListMessage() (interface{}, *errors.HTTPError)
	CreateHelpGuideMessage() (interface{}, *errors.HTTPError)
}

type SlashCommandRequest struct {
	UserID    string
	ChannelID string
	Command   consts.SlashCommand
	Arg       string
}
