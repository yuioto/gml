package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
	"github.com/yuioto/gml/core/downloader"
)

var XCmd = &cli.Command{
	Name:  "x",
	Usage: "Execute an minecraft version",
	Action: func(ctx context.Context, cli *cli.Command) error {
		configPath, err := os.UserConfigDir()
		if err != nil {
			return err
		}

		runPath := filepath.Join(configPath, "gml")
		if err := os.MkdirAll(runPath, 0755); err != nil {
			return err
		}
		if err := os.Chdir(runPath); err != nil {
			return err
		}

		version := cli.Args().First()
		log.Info().Str("version", version).Msg("Running version...")

		if err := downloader.DownloadVanilla(version); err != nil {
			log.Fatal().Err(err).Msg("downloading version failed")
		}

		// TODO: booting

		return nil
	},
}
