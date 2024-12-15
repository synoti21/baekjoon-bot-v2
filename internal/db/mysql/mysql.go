package mysql

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/schema"
)

type mysql struct {
	endpoint string
	platform config.Platform
}

var _ db.Interface = (*mysql)(nil)

func New(cfg *config.DatabaseClientConfig) db.Interface {
	return &mysql{
		endpoint: "",
		platform: cfg.Platform,
	}
}

func (m *mysql) AddUser(userID string, baekjoonID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) DeleteUser(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) ModifyBaekjoonID(userID string, baekjoonID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) SetDailyProbRecommendTime(userID string, time time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) UnsetDailyProbRecommendTime(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) FindUserWithID(userID string) (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) FindUsersWithDailyProbTime(time time.Time) (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) FindAllUser() (*schema.User, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) FindProbWithID(probID int) (*schema.BaekjoonProb, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mysql) AddFeedback(content string) error {
	panic("not implemented") // TODO: Implement
}
