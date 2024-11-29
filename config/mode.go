package config

type BotMode string

const (
	Socket  = BotMode("socket")
	Webhook = BotMode("webhook")
)
