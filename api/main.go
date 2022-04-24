package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Info().Msg("Start server")
	defer log.Info().Msg("Stop server")
	// TODO
}
