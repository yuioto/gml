package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var UpdateCmd = &cli.Command{
	Name:  "update",
	Usage: "Update a mod or resource",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "vanilla",
			Aliases: []string{"v"},
			Usage:   "The vanilla version to be updated",
		},
	},
	Action: func(ctx context.Context, cli *cli.Command) error {
		vanilla := cli.String("vanilla")
		if vanilla != "" {
			log.Info().Str("vanilla", vanilla).Msg("Updating vanilla version...")
		}
		log.Info().Msg("Updating mod or resource...")
		return nil
	},
}
