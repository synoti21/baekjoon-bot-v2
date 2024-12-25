package bots

import (
	"fmt"
	"strconv"
	"time"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/client"
	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Bot struct {
	db      db.Interface
	probRec client.ProblemRecommendClient
}

var _ Interface = (*Bot)(nil)

func New(_db db.Interface, _prc client.ProblemRecommendClient) Interface {
	return &Bot{
		db:      _db,
		probRec: _prc,
	}
}

func (b *Bot) RegisterUser(userID string, bojID string) *errors.BaseError {
	err := b.db.AddUser(userID, bojID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (b *Bot) WithdrawUser(userID string) *errors.BaseError {
	err := b.db.DeleteUser(userID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (b *Bot) GetRecommendedProb(userID string) (*schema.BaekjoonProb, *errors.BaseError) {
	resp, err := b.probRec.GetProblemsByUserID(userID, 1)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	pid, ok := resp.ProblemIDsByUserID[userID]
	if len(pid) != 1 || !ok {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Invalid baekjoon problem response: %v", resp.ProblemIDsByUserID))
	}

	prob, err := b.db.FindProbWithID(pid[0])
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return prob, nil
}

func (b *Bot) GetRecommendedProbByCategory(userID string, categoryType string) (*schema.BaekjoonProb, *errors.BaseError) {
	pc, err := consts.ValidateProbCategory(categoryType)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	resp, err := b.probRec.GetProblemsByCategory(pc)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	pid, ok := resp.ProblemIDsByCategory[pc]
	if len(pid) != 1 || !ok {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Invalid baekjoon prob response: %v", resp.ProblemIDsByCategory))
	}

	prob, err := b.db.FindProbWithID(pid[0])
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return prob, nil
}

func (b *Bot) GetSimilarProbByID(probID string, userID string) (*schema.BaekjoonProb, *errors.BaseError) {
	pid, err := strconv.ParseInt(probID, 0, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("problem ID should be number")
	}

	resp, err := b.probRec.GetSimilarProblemsByProblemIDs(int(pid))
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	pidResp, ok := resp.SimilarProblemIDsByProblemID[probID]
	if len(pidResp) != 1 || !ok {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Invalid baekjoon prob response: %v", resp.ProblemIDsByCategory))
	}

	prob, err := b.db.FindProbWithID(pidResp[0])
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return prob, nil
}

// This function is not implemented in current version
func (b *Bot) GetSimilarProbByContent(probContent string, userID string) (*schema.BaekjoonProb, *errors.BaseError) {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ScheduleDailyProbRecommend(userID string, _time string) *errors.BaseError {
	t, err := time.Parse("15 04", _time)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = b.db.SetDailyProbRecommendTime(userID, t)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (b *Bot) UnscheduleDailyProbRecommend(userID string) *errors.BaseError {
	err := b.db.UnsetDailyProbRecommendTime(userID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
