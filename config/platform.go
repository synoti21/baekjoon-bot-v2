package config

type Platform string

const (
	Slack   = Platform("slack")
	Discord = Platform("discord")
	Local   = Platform("local")
)
