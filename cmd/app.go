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
			Description: "This command fetches all metrics via Datadog API and categorizes with teams",
			Action:      cg.Start,
			Flags: []cli.Flag{
				&flagPort,
				&flagPlatform,
			},
		},
		&cli.Command{
			Name:        "daily",
			Description: "This command generates weekly usage data of custom metrics by team",
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

	cfg := config.New(platform, mode, port)

	switch cfg.Database() {
	case config.Mongo:
		db = mongo.New()
	default:
		return errors.NewInternalServerError("Invalid database mode")
	}

	api, err := client.NewProbRecommandSvc()
	if err != nil {
		return err
	}

	bot := bots.New(db, api)

	h, err := handlers.New(cfg, bot)
	if err != nil {
		return err
	}

	h.Run()

	return nil
}

func (cg *CommandGroup) Daily(cliCtx *cli.Context) error {
	panic("not implemented") //TODO: implement
}
