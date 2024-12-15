package cmd

import (
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
			Description: "Starts the Baekjoon bot server",
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
	var db db.Interface

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

	switch dcfg.Type {
	case config.DatabaseTypeMongoDB:
		db = mongo.New(dcfg)
	case config.DatabaseTypeMySQL:
		return errors.NewInternalServerError("MySQL not supported in this version")
	case config.DatabaseTypePostgres:
		return errors.NewInternalServerError("Postgres not supported in this version")
	case config.DatabaseTypeDryRun:
		return errors.NewInternalServerError("DryRun not supported in this version")
	default:
		return errors.NewInternalServerError("Invalid database mode")
	}

	recAPI, err := client.NewProbRecommandSvc()
	if err != nil {
		return err
	}

	bot := bots.New(db, recAPI)

	h, err := handlers.New(hcfg, bot)
	if err != nil {
		return err
	}
	h.Run()
	return nil
}

func (cg *CommandGroup) Daily(cliCtx *cli.Context) error {
	panic("not implemented") //TODO: implement
}
