package apiserver

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/blinnikov/go-rest-api/internal/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

func Start(config *Config, logger *logrus.Logger) error {
	address := strings.Split(config.DatabaseURL, " password")[0]
	logger.Printf("Starting apiserver with db %s", address)
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(logger, store, sessionStore)

	certFile := "certs/go-rest-api.crt"
	keyFile := "certs/go-rest-api.key"
	return http.ListenAndServeTLS(config.BindAddr, certFile, keyFile, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
