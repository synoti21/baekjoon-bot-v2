package config

import (
	"time"
)

const (
	defaultIdleTimeout = 20 * time.Second
)

type HandlerConfig struct {
	Port        string
	Secret      string
	Platform    Platform
	BotMode     BotMode
	IdleTimeout time.Duration
}

func NewHandlerConfig(platform Platform, botmode BotMode, port string) *HandlerConfig {
	return &HandlerConfig{
		Port:        port,
		Platform:    platform,
		BotMode:     botmode,
		IdleTimeout: defaultIdleTimeout,
	}
}

// RouteEndpoint returns endpoint that gets the slash command request, depending on the platform
// Slack receives the slash command request from '/receive'
// Discord receives the slash command request from '/interaction'
func (c *HandlerConfig) RouteEndpoint() string {
	switch c.Platform {
	case Slack:
		return "/receive"
	case Discord:
		return "/interaction"
	default:
		return "/default"
	}
}
