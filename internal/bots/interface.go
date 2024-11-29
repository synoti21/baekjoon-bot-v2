package bots

import (
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Interface interface {
	RegisterUser(userID, bojID string) *errors.HTTPError
	WithdrawUser(userID string) *errors.HTTPError

	GetRecommendedProb(userID string) (*schema.BaekjoonProb, *errors.HTTPError)
	GetRecommendedProbByCategory(userID string, categoryType string) (*schema.BaekjoonProb, *errors.HTTPError)
	GetSimilarProbByID(probID, userID string) (*schema.BaekjoonProb, *errors.HTTPError)
	GetSimilarProbByContent(probContent, userID string) (*schema.BaekjoonProb, *errors.HTTPError)

	ScheduleDailyProb(userID string, time string) *errors.HTTPError
	UnscheduleDailyProb(userID string) *errors.HTTPError

	ShowProbCategoryList(userID string) *errors.HTTPError
	ShowHelpGuide(userID string) *errors.HTTPError
}
