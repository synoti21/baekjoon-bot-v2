package handlers

import (
	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
)

func RunBotSlashCommand(b bots.Interface, req *adapters.SlashCommandRequest) (interface{}, *errors.HTTPError) {
	switch req.Command {
	case consts.SCRegister:
		return nil, b.RegisterUser(req.UserID, req.ChannelID)
	case consts.SCWithdraw:
		return nil, b.WithdrawUser(req.UserID)
	case consts.SCRecommandSingleProblem:
		return b.GetRecommendedProb(req.UserID)
	case consts.SCRecommandByCategory:
		return b.GetRecommendedProbByCategory(req.UserID, req.Arg)
	case consts.SCRecommandSimilarProblem:
		return b.GetSimilarProbByID(req.Arg, req.UserID)
	case consts.SCScheduleDailyProblem:
		return nil, b.ScheduleDailyProb(req.UserID, req.Arg)
	case consts.SCUnscheduleDailyProblem:
		return nil, b.UnscheduleDailyProb(req.UserID)
	case consts.SCShowCategoryList:
		return nil, b.ShowProbCategoryList(req.UserID)
	case consts.SCShowHelpGuide:
		return nil, b.ShowHelpGuide(req.UserID)
	default:
		return nil, errors.NewInvalidSlashCommandError("Unknown slash command")
	}
}
