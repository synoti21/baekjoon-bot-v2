package handlers

import (
	"fmt"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/adapters"
)

// RunBotSlashCommand runs the main logics depending on the slash command type.
func (h *SlashCommandHandler) RunBotSlashCommand(req *adapters.SlashCommandRequest) (interface{}, *errors.HTTPError) {
	switch req.Command {
	// /register
	case consts.SCRegister:
		if err := h.bot.RegisterUser(req.UserID, req.ChannelID); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return h.adapter.CreateTextMessage("성공적으로 등록했습니다.")
	// /quit
	case consts.SCWithdraw:
		if err := h.bot.WithdrawUser(req.UserID); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return h.adapter.CreateTextMessage("성공적으로 탈퇴했습니다.")
	// /prob
	case consts.SCRecommendSingleProblem:
		p, err := h.bot.GetRecommendedProb(req.UserID)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	// /category
	case consts.SCRecommendByCategory:
		p, err := h.bot.GetRecommendedProbByCategory(req.UserID, req.Arg)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	// /similarid
	case consts.SCRecommendSimilarProblem:
		p, err := h.bot.GetSimilarProbByID(req.Arg, req.UserID)
		if err != nil {
			return nil, err
		}
		return h.adapter.CreateProblemMessage(p)
	// /daily
	case consts.SCScheduleDailyProblem:
		if err := h.bot.ScheduleDailyProbRecommend(req.UserID, req.Arg); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return h.adapter.CreateTextMessage(fmt.Sprintf("성공적으로 일일 문제 알림을 등록했습니다. 시간: %s", req.Arg))
	// /deactivate
	case consts.SCUnscheduleDailyProblem:
		if err := h.bot.UnscheduleDailyProbRecommend(req.UserID); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return h.adapter.CreateTextMessage("성공적으로 일일 문제 알림을 해제했습니다.")
	// /categorylist
	case consts.SCShowCategoryList:
		return h.adapter.CreateCategoryListMessage()
	// /help
	case consts.SCShowHelpGuide:
		return h.adapter.CreateHelpGuideMessage()
	default:
		return nil, errors.NewInvalidSlashCommandError("Unknown slash command")
	}
}
