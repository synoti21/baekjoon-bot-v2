package bots

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
)

type Interface interface {
	RegisterUser(userID, chanID string) error
	WithdrawUser(userID, chanID string) error

	SendProbToUser(userID string) error
	SendProbToUserByCategory(userID, chanID string, pc consts.ProbCategory) error
	SendSimliarProbByPID(probID, userID, chanID string) error
	SendSimilarProbByContent(probContent, userID, chanID string) error

	ScheduleDailyProb(userID, chanID string, time time.Time) error
	UnscheduleDailyProb(userID, chanID string) error

	ShowProbCategoryList(userID, chanID string) error
	ShowHelpGuide(userID, chanID string) error
}
