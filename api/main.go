package main

import (
	"api/interface/handler"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Info().Msg("Start server")
	defer log.Info().Msg("Stop server")
	// Router
	e := echo.New()
	e.GET("/", handler.Health)
	log.Fatal().Err(e.Start(":80"))
}
