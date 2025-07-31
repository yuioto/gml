package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "Run the project",
	Action: func(ctx context.Context, cli *cli.Command) error {
		path := cli.Args().First()
		log.Info().Str("path", path).Msg("Running project...")
		return nil
	},
}
