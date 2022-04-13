package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("certtools failed to run")
	}
}
