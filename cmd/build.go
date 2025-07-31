package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var BuildCmd = &cli.Command{
	Name:  "build",
	Usage: "Build the project",
	Action: func(context.Context, *cli.Command) error {
		log.Info().Msg("Building the project...")
		return nil
	},
}
