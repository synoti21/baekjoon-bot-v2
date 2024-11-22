package bots

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
)

type Interface interface {
	Init() error

	RegisterUser(uid string) error
	WithdrawUser(uid string) error

	SendProbToUser(uid string) error
	SendProbToUserByCategory(uid string, pc consts.ProbCategory) error
	SendSimliarProbByPID(pid string, uid string) error
	SendSimilarProbByContent(pctnt string, uid string) error

	ScheduleDailyProb(uid string, time time.Time) error
	UnscheduleDailyProb(uid string) error

	ShowProbCategoryList(uid string) error
	ShowHelpGuide(uid string) error
}
