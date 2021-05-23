package store

import (
	"database/sql"
	"log"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
