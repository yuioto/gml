package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var ImportCmd = &cli.Command{
	Name:  "import",
	Usage: "Import integration package",
	Action: func(ctx context.Context, cli *cli.Command) error {
		path := cli.Args().First()
		log.Info().Str("path", path).Msg("Importing integration package...")
		return nil
	},
}
