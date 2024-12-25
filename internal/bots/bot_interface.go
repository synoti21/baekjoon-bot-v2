package bots

import (
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

// Interface is an interface of bot, which performs the main logics of each command
type Interface interface {
	// RegisterUser handles /register command, to add user information mapped with the userID.
	// We have to distinguish the slack ID and discord ID (or other possible platforms), which can be achieved by `Platform`
	RegisterUser(userID, bojID string) *errors.BaseError
	// WithdrawUser handles /quit command, to delete the user information.
	WithdrawUser(userID string) *errors.BaseError

	// GetRecommendedProb handles /prob command, giving user a recommended problem.
	GetRecommendedProb(userID string) (*schema.BaekjoonProb, *errors.BaseError)
	// GetRecommendedProbByCategory handles /category command, giving user a recommend problem of the specific category.
	GetRecommendedProbByCategory(userID string, categoryType string) (*schema.BaekjoonProb, *errors.BaseError)
	// GetSimilarProbByID handles /similarid command, giving user a similar one of the given problem with the id.
	GetSimilarProbByID(probID, userID string) (*schema.BaekjoonProb, *errors.BaseError)
	// GetSimilarProbByContent handles /similarprob command, giving user a similar one of the given problem content.
	// NOTE: This command will be implemented in the future.
	GetSimilarProbByContent(probContent, userID string) (*schema.BaekjoonProb, *errors.BaseError)

	// ScheduleDailyProbRecommend handles /daily command, setting the daily problem recommendation time.
	ScheduleDailyProbRecommend(userID string, time string) *errors.BaseError
	// UnscheduleDailyProbRecommend handles /deactivate command, unsetting the daily problem recommendation time.
	UnscheduleDailyProbRecommend(userID string) *errors.BaseError
}
