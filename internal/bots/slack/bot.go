package slack

import (
	"time"

	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
	"github.com/synoti21/baekjoon-slack-bot/internal/common/utils"
)

type Bot struct {
	botToken string
}

var _ bots.Interface = (*Bot)(nil)

func NewSlackBot() bots.Interface {
	t, err := utils.GetVarFromEnv("SLACK_BOT_TOKEN")
	if err != nil {
		return nil
	}
	return &Bot{
		botToken: t,
	}
}

func (b *Bot) RegisterUser(userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) WithdrawUser(userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendProbToUser(userID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendProbToUserByCategory(userID string, chanID string, pc consts.ProbCategory) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendSimliarProbByPID(probID string, userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendSimilarProbByContent(probContent string, userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ScheduleDailyProb(userID string, chanID string, time time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) UnscheduleDailyProb(userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowProbCategoryList(userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowHelpGuide(userID string, chanID string) error {
	panic("not implemented") // TODO: Implement
}
