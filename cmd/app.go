package cmd

import (
	"fmt"

	"github.com/synoti21/baekjoon-slack-bot/api/handlers"
	"github.com/synoti21/baekjoon-slack-bot/common/errors"
	"github.com/synoti21/baekjoon-slack-bot/config"
	"github.com/synoti21/baekjoon-slack-bot/internal/bots"
	"github.com/synoti21/baekjoon-slack-bot/internal/client"
	"github.com/synoti21/baekjoon-slack-bot/internal/db"
	"github.com/synoti21/baekjoon-slack-bot/internal/db/mongo"
	"github.com/urfave/cli/v2"
)

type CommandGroup struct{}

func (cg *CommandGroup) RegisterApp(app *cli.App) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:        "start",
			Description: "Starts the Baekjoon bot server for Slack or Discord. Example: baekjoon-bot start --platform slack --mode proxy --port 8080",
			Action:      cg.Start,
			Flags: []cli.Flag{
				&flagPort,
				&flagBotMode,
				&flagPlatform,
			},
		},
		&cli.Command{
			Name:        "daily",
			Description: "Schedules daily problem recommendations",
			Action:      cg.Daily,
			Flags: []cli.Flag{
				&flagPlatform,
			},
		},
	)
}

func (cg *CommandGroup) Start(cliCtx *cli.Context) error {
	platform, err := parsePlatformFromCtx(cliCtx)
	if err != nil {
		return err
	}
	port, err := parsePortFromCtx(cliCtx)
	if err != nil {
		return err
	}
	mode, err := parseBotModeFromCtx(cliCtx)
	if err != nil {
		return err
	}

	hcfg := config.NewHandlerConfig(platform, mode, port)
	dcfg := config.NewDatabaseClientConfig()
	if err = dcfg.Validate(); err != nil {
		return err
	}

	var db db.Interface
	switch dcfg.Type {
	case config.DatabaseTypeMongoDB:
		db = mongo.New(dcfg)
	case config.DatabaseTypeMySQL, config.DatabaseTypePostgres, config.DatabaseTypeDryRun:
		return errors.NewInternalServerError(fmt.Sprintf("Database type %s not supported in this version", dcfg.Type))
	default:
		return errors.NewInternalServerError("Invalid database mode")
	}

	pr, err := client.NewProblemRecommendClient()
	if err != nil {
		return err
	}

	bot := bots.New(db, pr)

	h, err := handlers.New(hcfg, bot)
	if err != nil {
		return err
	}

	return h.Run()
}

func (cg *CommandGroup) Daily(cliCtx *cli.Context) error {
	return errors.NewInternalServerError("Daily command is not implemented yet")
}
