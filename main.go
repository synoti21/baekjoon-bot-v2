package main

import (
	"os"

	"github.com/synoti21/baekjoon-slack-bot/cmd"
	"github.com/synoti21/baekjoon-slack-bot/common/tools/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "baekjoon-bot",
		Description: "Baekjoon Bot for recommanding algorithm problems",
		Usage:       "baekjoon-bot <command> [OPTIONS]",
		Version:     "1.0.0",
	}
	cmdGroup := new(cmd.CommandGroup)
	cmdGroup.RegisterApp(app)

	sugar := logger.GetLogger().Sugar()
	if err := app.Run(os.Args); err != nil {
		sugar.Fatal(err)
	}
}
