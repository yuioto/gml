package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var CheckCmd = &cli.Command{
	Name:  "check",
	Usage: "Check integrity",
	Action: func(context.Context, *cli.Command) error {
		log.Info().Msg("Checking integrity...")
		return nil
	},
}
