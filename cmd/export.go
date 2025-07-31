package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var ExportCmd = &cli.Command{
	Name:  "export",
	Usage: "Export integration package",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:             "format",
			Aliases:          []string{"f"},
			Usage:            "Export format (gmlpack, mrpack, etc.)",
			Value:            "gmlpack",
			ValidateDefaults: true,
		},
		&cli.StringFlag{
			Name:             "output",
			Aliases:          []string{"o"},
			Usage:            "Output file path",
			Value:            "output.gmlpack",
			TakesFile:        true,
			ValidateDefaults: true,
		},
	},
	Action: func(ctx context.Context, cli *cli.Command) error {
		format := cli.String("format")
		output := cli.String("output")
		log.Info().Str("format", format).Str("output", output).Msg("Exporting integration package...")
		return nil
	},
}
