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

func (b *Bot) Init() error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) RegisterUser(uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) WithdrawUser(uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendProbToUser(uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendProbToUserByCategory(uid string, pc consts.ProbCategory) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendSimliarProbByPID(pid string, uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) SendSimilarProbByContent(pctnt string, uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ScheduleDailyProb(uid string, time time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) UnscheduleDailyProb(uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowProbCategoryList(uid string) error {
	panic("not implemented") // TODO: Implement
}

func (b *Bot) ShowHelpGuide(uid string) error {
	panic("not implemented") // TODO: Implement
}
