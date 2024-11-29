package config

type Platform string

const (
	Slack   = Platform("slack")
	Discord = Platform("discord")
	Test    = Platform("test")
)
