package bots

import (
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/client"
	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Bot struct {
	db     db.Interface
	recAPI client.ProbRecommandAPI
}

var _ Interface = (*Bot)(nil)

func (b *Bot) RegisterUser(userID string, bojID string) *errors.HTTPError {
	err := b.db.AddUser(userID, bojID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (b *Bot) WithdrawUser(userID string) *errors.HTTPError {
	err := b.db.DeleteUser(userID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (b *Bot) GetRecommendedProb(userID string) (*schema.BaekjoonProb, *errors.HTTPError) {
	p, err := b.recAPI.GetProbsByUserID(userID, 1)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	pid := p.ProbIDsByUserID[userID]
	if len(pid) != 1 {
		return nil, errors.NewInternalServerError("Invalid response from prob recommand service")
	}

	prob, err := b.db.FindProbWithID(pid[0])
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return prob, nil
}

func (b *Bot) GetRecommendedProbByCategory(userID string, categoryType string) (*schema.BaekjoonProb, *errors.HTTPError) {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) GetSimilarProbByID(probID string, userID string) (*schema.BaekjoonProb, *errors.HTTPError) {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) GetSimilarProbByContent(probContent string, userID string) (*schema.BaekjoonProb, *errors.HTTPError) {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ScheduleDailyProb(userID string, time string) *errors.HTTPError {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) UnscheduleDailyProb(userID string) *errors.HTTPError {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowProbCategoryList(userID string) *errors.HTTPError {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowHelpGuide(userID string) *errors.HTTPError {
	panic("not implemented") // TODO: Implement
}
