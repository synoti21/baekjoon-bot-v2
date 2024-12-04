package db

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type Interface interface {
	AddUser(userID, bojID string) error
	DeleteUser(userID string) error
	ModifyUserBJID(userID, BJID string) error

	SetDailyProbTime(userID string, time time.Time) error
	UnsetDailyProbTime(userID string) error

	FindUserWithDiscordID(discordID string)
	FindUserWithSlackID(slackID string)
	FindUsersWithDailyProbTime(time time.Time)
	FindAllUser()

	FindProbWithID(probID int) (*schema.BaekjoonProb, error)

	AddFeedback(content string)
}
