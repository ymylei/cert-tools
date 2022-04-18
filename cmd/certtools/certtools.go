package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/ymylei/cert-tools/cmd/certtools/generate"
)

func main() {
	app := &cli.App{
		HelpName: "Random tools for cert generation and verification/validation",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generates a self-signed cert and associated private key.",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args()
					if args.Present() {
						err := generate.Generate(args.First())
						if err != nil {
							log.Fatal().Err(err).Str("name", args.First()).Msg("failed to generate random cert")
						}
					}
					log.Warn().Msg("No args passed, exiting")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("certtools failed to run")
	}
}
