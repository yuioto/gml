package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var NewCmd = &cli.Command{
	Name:  "new",
	Usage: "Create a new project",
	Action: func(ctx context.Context, cli *cli.Command) error {
		path := cli.Args().First()
		log.Info().Str("path", path).Msg("Creating new project...")
		return nil
	},
}
