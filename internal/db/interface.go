package db

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Interface interface {
	AddUser(userID, baekjoonID string) error
	DeleteUser(userID string) error
	ModifyBaekjoonID(userID, baekjoonID string) error

	SetDailyProbRecommendTime(userID string, time time.Time) error
	UnsetDailyProbRecommendTime(userID string) error

	FindUserWithID(userID string) (*schema.User, error)
	FindUsersWithDailyProbTime(time time.Time) (*schema.User, error)
	FindAllUser() (*schema.User, error)

	FindProbWithID(probID int) (*schema.BaekjoonProb, error)

	AddFeedback(content string) error
}
