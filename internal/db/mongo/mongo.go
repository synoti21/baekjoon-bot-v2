package mongo

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type mongo struct {
	endpoint string
	platform config.Platform
}

var _ db.Interface = (*mongo)(nil)

func New(cfg *config.DatabaseClientConfig) db.Interface {
	return &mongo{
		endpoint: "",
		platform: cfg.Platform,
	}
}

func (m *mongo) AddUser(userID string, bojID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) DeleteUser(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) ModifyBaekjoonID(userID string, BJID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) SetDailyProbRecommendTime(userID string, time time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) UnsetDailyProbRecommendTime(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindUserWithID(userID string) (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindUsersWithDailyProbTime(time time.Time) (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindAllUser() (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) FindProbWithID(probID int) (*schema.BaekjoonProb, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mongo) AddFeedback(content string) error {
	panic("not implemented") // TODO: Implement
}
