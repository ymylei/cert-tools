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
				Usage: "generates various test objects",
				Subcommands: []*cli.Command{
					{
						Name:  "cert",
						Usage: "generates testing cert",
						Action: func(ctx *cli.Context) error {
							args := ctx.Args()
							if args.Present() {
								err := generate.Generate(args.First())
								if err != nil {
									log.Fatal().Err(err).Str("name", args.First()).Msg("failed to generate random cert")
								}
							} else {
								log.Warn().Msg("No args passed, exiting")
							}
							return nil
						},
					},
					{
						Name:  "keypair",
						Usage: "generates a ECDSA key pair and outputs the priv/pub keys as pem files.",
						Action: func(ctx *cli.Context) error {
							args := ctx.Args()
							if args.Present() {
								err := generate.GenerateKeyPair(args.First())
								if err != nil {
									log.Fatal().Err(err).Str("name", args.First()).Msg("failed to generate key pair")
								}
							} else {
								log.Warn().Msg("No args passed, exiting")
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("certtools failed to run")
	}
}
