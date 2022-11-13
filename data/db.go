package data

import (
	"log"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

var DbSession *Database

// Database is a construct of bond.Session and bond stores
type Database struct {
	db.Session
}

func NewDBSession() (*Database, error) {
	settings := sqlite.ConnectionURL{
		Database: "krek.db",
	}

	// Attempt to open the 'krek.db' database file
	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	DbSession = &Database{Session: sess}

	return DbSession, nil
}
