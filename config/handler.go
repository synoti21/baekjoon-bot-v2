package config

import (
	"os"
	"time"
)

const (
	defaultIdleTimeout = 20 * time.Second
)

type HandlerConfig struct {
	Secret      string
	IdleTimeout time.Duration
}

func New(secret string) *HandlerConfig {
	return &HandlerConfig{
		Secret:      secret,
		IdleTimeout: defaultIdleTimeout,
	}

}

func (c *HandlerConfig) Database() Database {
	db := os.Getenv("DATABASE_MODE")
	if db == "" {
		panic("DATABASE_MODE not set")
	}
	switch db {
	case "MONGO":
		return Mongo
	case "LOCAL":
		return Local
	default:
		panic("Invalid database mode")
	}
}

func (c *HandlerConfig) Platform() Platform {
	plat := os.Getenv("BOT_PLATFORM")
	if plat == "" {
		panic("BOT_PLATFORM not set")
	}
	switch plat {
	case "SLACK":
		return Slack
	case "DISCORD":
		return Discord
	default:
		panic("Invalid platform")
	}
}

func (c *HandlerConfig) RouteEndpoint() string {
	switch c.Platform() {
	case Slack:
		return "/receive"
	case Discord:
		return "/interaction"
	default:
		return "/"
	}
}

func (c *HandlerConfig) BotMode() BotMode {
	mode := os.Getenv("BOT_MODE")
	if mode == "" {
		panic("BOT_MODE not set")
	}
	switch mode {
	case "SOCKET":
		return Socket
	case "WEBHOOK":
		return Webhook
	default:
		panic("Unknown mode")
	}
}
