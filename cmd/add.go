package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var AddCmd = &cli.Command{
	Name:  "add",
	Usage: "Add a mod or resource",
	Commands: []*cli.Command{
		addModCmd,
		addResourceCmd,
		addDatapackCmd,
	},
	Action: func(ctx context.Context, cli *cli.Command) error {
		return addModCmd.Action(ctx, cli)
	},
}

var addModCmd = &cli.Command{
	Name:  "mod",
	Usage: "Add a mod",
	Action: func(ctx context.Context, cli *cli.Command) error {
		modName := cli.Args().First()
		log.Info().Str("mod", modName).Msg("Adding mod...")
		return nil
	},
}

var addResourceCmd = &cli.Command{
	Name:  "resource",
	Usage: "Add a resource",
	Action: func(ctx context.Context, cli *cli.Command) error {
		resourceName := cli.Args().First()
		log.Info().Str("resource", resourceName).Msg("Adding resource...")
		return nil
	},
}

var addDatapackCmd = &cli.Command{
	Name:  "datapack",
	Usage: "Add a datapack",
	Action: func(ctx context.Context, cli *cli.Command) error {
		datapackName := cli.Args().First()
		log.Info().Str("datapack", datapackName).Msg("Adding datapack...")
		return nil
	},
}
