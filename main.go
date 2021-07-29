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
	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
		NoColor:    true,
		FormatLevel: func(level interface{}) string {
			return fmt.Sprintf("%v", level)
		},
	}
	log.Logger = log.Output(writer)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(Logger(), gin.Recovery())
	delayResponse := os.Getenv("DELAY_RESPONSE")
	if delayResponse == "" {
		delayResponse = "10s"
	}
	duration, err := time.ParseDuration(delayResponse)
	if err != nil {
		log.Fatal().Err(err).Msg("DELAY_RESPONSE parsing failed")
	}
	r.NoRoute(func(context *gin.Context) {
		log.Info().Msgf("Sleeping for %v", duration)
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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		if query != "" {
			path += "?" + query
		}

		log.Info().Timestamp().
			Str("clientIp", c.ClientIP()).
			Str("path", path).
			Str("method", c.Request.Method).
			Msg("Received request")

		c.Next()

		log.Info().Dur("latency", time.Now().Sub(start)).
			Str("clientIp", c.ClientIP()).
			Str("path", path).
			Str("method", c.Request.Method).
			Int("statusCode", c.Writer.Status()).
			Msg("Sent response")
	}
}
