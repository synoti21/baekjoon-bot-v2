package mongo

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type mongo struct {
	endpoint string
}

var _ db.Interface = (*mongo)(nil)

func New() db.Interface {
	return &mongo{
		endpoint: "",
	}
}

func (m *mongo) AddUser(userID string, bojID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) DeleteUser(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) ModifyUserBJID(userID string, BJID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) UpsertDailyProbTime(userID string, time time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) UnsetDailyProbTime(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindUserWithDiscordID(discordID string) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindUserWithSlackID(slackID string) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindUsersWithDailyProbTime(time time.Time) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindAllUser() {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindProbWithID(probID int) (*schema.BaekjoonProb, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) AddFeedback(content string) {
	panic("not implemented") // TODO: Implement
}
