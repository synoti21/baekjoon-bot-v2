package db

import "time"

type Interface interface {
	AddUser(userID, bojID string)
	DeleteUser(userID string)
	ModifyUserBJID(userID, BJID string)

	UpsertDailyProbTime(userID string, time time.Time)
	UnsetDailyProbTime(userID string)

	FindUserWithDiscordID(discordID string)
	FindUserWithSlackID(slackID string)
	FindUsersWithDailyProbTime(time time.Time)
	FindAllUser()

	FindProbWithID(probID string)

	AddFeedback(content string)
}
