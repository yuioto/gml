package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var SearchCmd = &cli.Command{
	Name:  "search",
	Usage: "Search for a mod or vanilla",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "vanilla",
			Aliases: []string{"v"},
			Usage:   "Vanilla name to search",
		},
		&cli.BoolFlag{
			Name:    "mod",
			Aliases: []string{"m"},
			Usage:   "Mod name to search",
		},
		&cli.BoolFlag{
			Name:    "resource",
			Aliases: []string{"r"},
			Usage:   "Resource name to search",
		},
		&cli.BoolFlag{
			Name:    "datapack",
			Aliases: []string{"d"},
			Usage:   "Datapack name to search",
		},
	},
	Action: func(ctx context.Context, cli *cli.Command) error {
		vanilla := cli.Bool("vanilla")
		mod := cli.Bool("mod")
		resource := cli.Bool("resource")
		datapack := cli.Bool("datapack")

		if !vanilla && !mod && !resource && !datapack {
			log.Debug().
				Bool("vanilla", vanilla).
				Bool("mod", mod).
				Bool("resource", resource).
				Bool("datapack", datapack).
				Msg("No type specified. Defaulting to --mod.")
			mod = true
		}

		log.Info().
			Bool("vanilla", vanilla).
			Bool("mod", mod).
			Bool("resource", resource).
			Bool("datapack", datapack).
			Msg("Searching for mod or vanilla...")

		return nil
	},
}
