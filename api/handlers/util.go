package handlers

import (
	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
)

func (h *SlashCommandHandler) RunBotSlashCommand(req *adapters.SlashCommandRequest) (interface{}, *errors.HTTPError) {
	switch req.Command {
	case consts.SCRegister:
		return nil, h.bot.RegisterUser(req.UserID, req.ChannelID)
	case consts.SCWithdraw:
		return nil, h.bot.WithdrawUser(req.UserID)
	case consts.SCRecommendSingleProblem:
		p, err := h.bot.GetRecommendedProb(req.UserID)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	case consts.SCRecommendByCategory:
		p, err := h.bot.GetRecommendedProbByCategory(req.UserID, req.Arg)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	case consts.SCRecommendSimilarProblem:
		p, err := h.bot.GetSimilarProbByID(req.Arg, req.UserID)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	case consts.SCScheduleDailyProblem:
		return nil, h.bot.ScheduleDailyProb(req.UserID, req.Arg)
	case consts.SCUnscheduleDailyProblem:
		return nil, h.bot.UnscheduleDailyProb(req.UserID)
	case consts.SCShowCategoryList:
		return h.adapter.CreateCategoryListMessage()
	case consts.SCShowHelpGuide:
		return h.adapter.CreateHelpGuideMessage()
	default:
		return nil, errors.NewInvalidSlashCommandError("Unknown slash command")
	}
}
