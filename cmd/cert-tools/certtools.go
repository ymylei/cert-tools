package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		HelpName: "Random tools for cert generation and verification/validation",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generates a self-signed cert and associated private key.",
				Action: func(ctx *cli.Context) error {
					fmt.Printf("not implemented")
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
