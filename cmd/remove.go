package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var RemoveCmd = &cli.Command{
	Name:  "remove",
	Usage: "Remove a mod or resource",
	Commands: []*cli.Command{
		removeModCmd,
		removeResourceCmd,
		removeDatapackCmd,
	},
	Action: func(ctx context.Context, cli *cli.Command) error {
		return removeModCmd.Action(ctx, cli)
	},
}

var removeModCmd = &cli.Command{
	Name:  "mod",
	Usage: "Remove a mod",
	Action: func(ctx context.Context, cli *cli.Command) error {
		modName := cli.Args().First()
		log.Info().Str("mod", modName).Msg("Removing mod...")
		return nil
	},
}

var removeResourceCmd = &cli.Command{
	Name:  "resource",
	Usage: "Remove a resource",
	Action: func(ctx context.Context, cli *cli.Command) error {
		resourceName := cli.Args().First()
		log.Info().Str("resource", resourceName).Msg("Removing resource...")
		return nil
	},
}

var removeDatapackCmd = &cli.Command{
	Name:  "datapack",
	Usage: "Remove a datapack",
	Action: func(ctx context.Context, cli *cli.Command) error {
		datapackName := cli.Args().First()
		log.Info().Str("datapack", datapackName).Msg("Removing datapack...")
		return nil
	},
}
