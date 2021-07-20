package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano, NoColor: true, FormatLevel: func(level interface{}) string {
		return fmt.Sprintf("%v", level)
	}})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	delayResponse := os.Getenv("DELAY_RESPONSE")
	if delayResponse == "" {
		delayResponse = "10s"
	}
	duration, err := time.ParseDuration(delayResponse)
	if err != nil {
		log.Fatal().Err(err).Msg("DELAY_RESPONSE parsing failed")
	}
	r.NoRoute(func(context *gin.Context) {
		request := context.Request
		log.Info().
			Str("url", request.URL.String()).Str("method", request.Method).Msg("Received request")
		log.Printf("Received request method. Sleeping for %v", duration)
		time.Sleep(duration)
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	log.Info().Msg("Starting server on port 8080")
	err = r.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}