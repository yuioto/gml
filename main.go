package main

import (
	"context"
	"os"
	"time"

	"github.com/yuioto/gml/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	app := &cli.Command{
		Name:                  "gml",
		Usage:                 "Manage Minecraft launchers and instances",
		Version:               "0.1.0",
		Commands:              cmd.Commands,
		EnableShellCompletion: true,
		Suggest:               true,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("Application failed to start")
	}
}
