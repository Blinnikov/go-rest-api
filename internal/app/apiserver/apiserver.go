package apiserver

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/blinnikov/go-rest-api/internal/bus"
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

	go messageReceiver(logger)
	go messageSender(logger)

	certFile := "certs/go-rest-api.crt"
	keyFile := "certs/go-rest-api.key"
	return http.ListenAndServeTLS(config.BindAddr, certFile, keyFile, srv)
}

func messageSender(logger *logrus.Logger) {
	conn, ch, q := bus.GetQueue()
	defer conn.Close()
	defer ch.Close()

	for {
		msg := "Golang is awesome - Keep Moving Forward!"
		bus.SendTextMessage(ch, q.Name, msg)
		// logger.Printf(" [x] Congrats, sending message: %s", msg)
	}
}

func messageReceiver(logger *logrus.Logger) {
	conn, ch, q := bus.GetQueue()
	defer conn.Close()
	defer ch.Close()

	msgs := bus.ReceiveMessage(ch, q.Name)
	for msg := range msgs {
		logger.Printf("Received message from queue: %s", msg)
	}
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
