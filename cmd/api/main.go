package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LukasJenicek/jsi-krek/data"
	"github.com/LukasJenicek/jsi-krek/services/api"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/upper/db/v4"
)

func main() {
	addr := ":9099"

	server := http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	configureLogger()

	zlog.Info().Msgf("serving on port %s", addr)

	sess, err := data.NewDBSession()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	var questions []*data.SurveyQuestion
	err = data.GetSurveyQuestionStore(sess).Find(db.Cond{"survey_id": 1}).All(&questions)
	if err != nil {
		log.Fatal(err)
	}

	for _, question := range questions {
		zlog.Info().Msgf("%d: %s", question.Id, question.Question)
	}

	err = server.ListenAndServe()
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
	log.SetFlags(0)
	log.SetOutput(logger)
	log.SetPrefix("")

	zlog.Logger = logger
}
