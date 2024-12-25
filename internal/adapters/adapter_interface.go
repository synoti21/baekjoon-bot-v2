package adapters

import (
	"net/http"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

// SlashCommnadRequest is a struct to handle slash command request without considering the platforms.
// Request JSON of slash command request differs by platform, so we use this struct to create a common
// instance of slash command request
type SlashCommandRequest struct {
	UserID    string
	ChannelID string
	Command   consts.SlashCommand
	// Arg receives an extra argument that user inputs when using a slash command.
	// Arg can be userID (/register), category number (/category) or other else.
	Arg string
}

// Interface is an interface of adapter, used to perform common logics in various platform like Slack, Discord.
// Since logics of Baekjoon-bot do not significantly differ by platforms, we use adapters to adaptively implement logics.
// It will be implemented with real logics depending on the platform we use.
// For example, Slack will verify slash commnad request with signature, while Discord verifies with token.
type Interface interface {
	// VerifyRequest verfies the slash command request from platforms, by verifying secret of header.
	// Unverified request will be aborted.
	VerifyRequest(r *http.Request, secret string) *errors.BaseError
	// ParseSlashCommnad will parse the slash command from request, and run the following command
	ParseSlashCommand(r *http.Request) (*SlashCommandRequest, *errors.BaseError)
	// CreateTextMessage creates a simple text message for either Slack or Discord
	CreateTextMessage(text string) (interface{}, *errors.BaseError)
	// CreateProblemMessage creates a message that will be sent to platforms like Slack or Discord.
	// Slack uses BlockMessage to send message, while Discord uses EmbedMessage.
	// These structures are different from each other, so we should use this function adaptively.
	CreateProblemMessage(prob *schema.BaekjoonProb) (interface{}, *errors.BaseError)
	// CreateCategoryListMessage creates a message that shows a category list of baekjoon problem.
	// Same as the reason above, we use this function to send messages depending on the platforms we use
	CreateCategoryListMessage() (interface{}, *errors.BaseError)
	// CreateHelpGuideMessage creates a message that shows a help guide including command list.
	// Same as the reason above.
	CreateHelpGuideMessage() (interface{}, *errors.BaseError)
}
