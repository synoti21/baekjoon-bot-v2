package client

import (
	"github.com/synoti21/baekjoon-slack-bot/internal/types/discord"
)

type Discord interface {
	GetUserID(user string) (string, error)
	GetChannelID(channel string) (string, error)
	SendMessage(msg discord.Embed, ID string) error
}
