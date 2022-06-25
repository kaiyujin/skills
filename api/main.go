package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Info().Msg("Start server")
	defer log.Info().Msg("Stop server")
	// Router
	e := echo.New()
	e.GET("/", health)
	log.Fatal().Err(e.Start(":80"))
}

func health(c echo.Context) error {
	result := map[string]interface{}{
		"status": "OK",
	}
	return c.JSON(http.StatusOK, result)
}
