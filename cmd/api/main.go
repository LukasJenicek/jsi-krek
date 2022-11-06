package main

import (
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/LukasJenicek/jsi-krek/services/api"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	addr := ":9099"

	server := http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	configureLogger()

	zlog.Info().Msgf("serving on port %s", addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TODO: Move to setup functionality
func configureLogger() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Caller().
		Timestamp().
		Logger()

	loc := time.Local
	zerolog.TimestampFunc = func() time.Time { return time.Now().In(loc) }
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = "15:04:05"
	zerolog.CallerSkipFrameCount = 3

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Redirect stdlib's log output to zerolog.
	stdlog.SetFlags(0)
	stdlog.SetOutput(logger)
	stdlog.SetPrefix("")

	zlog.Logger = logger
}
