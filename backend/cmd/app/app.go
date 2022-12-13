package app

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	flags := []cli.Flag{
		config.EnvFlag,
		config.AppNameFlag,
		config.AppPortFlag,
		config.PGSQLHostFlag,
		config.PGSQLPortFlag,
		config.PGSQLUsernameFlag,
		config.PGSQLPasswordFlag,
		config.PGSQLNameFlag,
		config.PGSQLMaxOpenConnsFlag,
		config.PGSQLMaxIdleConnsFlag,
		config.PGSQLConnMaxLifetimeFlag,
		config.PGSQLIsEnabledLogFlag,
		config.BasicAuthUsernameFlag,
		config.BasicAuthPasswordFlag,
	}

	app := &cli.App{
		Name:  "Cubicasa service",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			start(ctx)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
