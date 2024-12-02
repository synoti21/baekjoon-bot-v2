package cmd

import (
	"fmt"
	"strconv"

	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/urfave/cli/v2"
)

const (
	flagPortName     = "port"
	flagBotModeName  = "mode"
	flagPlatformName = "platform"
)

var (
	flagPlatform = cli.StringFlag{
		Name:     flagPlatformName,
		Usage:    "Set the platform to host baekjoon bot.",
		Required: true,
	}

	flagBotMode = cli.StringFlag{
		Name:     flagBotModeName,
		Usage:    "Set the bot mode (proxy / websocket)",
		Required: true,
	}

	flagPort = cli.StringFlag{
		Name:  flagPortName,
		Usage: "Set the port number",
	}
)

func parsePlatformFromCtx(cliCtx *cli.Context) (config.Platform, error) {
	platform := cliCtx.String(flagPlatformName)
	if platform == "" {
		return "", fmt.Errorf("no platform inputted")
	}
	switch platform {
	case "slack":
		return config.Slack, nil
	case "discord":
		return config.Discord, nil
	default:
		return "", errors.NewBadRequestError("Invalid platform")
	}
}

func parsePortFromCtx(cliCtx *cli.Context) (string, error) {
	port := cliCtx.String(flagPortName)
	if port == "" {
		return "", errors.NewBadRequestError("No port inputted")
	}

	if _, err := strconv.ParseInt(port, 0, 64); err != nil {
		return "", errors.NewBadRequestError("Invalid port number")
	}
	return port, nil
}

func parseBotModeFromCtx(cliCtx *cli.Context) (config.BotMode, error) {
	mode := cliCtx.String(flagBotModeName)
	if mode == "" {
		return "", errors.NewBadRequestError("No bot mode inputted")
	}

	switch mode {
	case "proxy":
		return config.Webhook, nil
	case "websocket":
		return config.Socket, nil
	default:
		return "", errors.NewBadRequestError("Invalid botmode")
	}
}
