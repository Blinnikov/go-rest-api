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

// TODO: Get from configuration
const addr string = "amqp://rabbitmq:rabbitmq@rabbit1:5672/"
const queueName string = "go-reat-api-queue"

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

	// Create queue on app start
	conn, ch := bus.GetChannel(addr)
	defer conn.Close()
	defer ch.Close()
	bus.GetQueue(queueName, ch)

	// Start reading/publishing messages
	go messageReceiver(logger)
	go messageSender(logger)

	certFile := "certs/go-rest-api.crt"
	keyFile := "certs/go-rest-api.key"
	return http.ListenAndServeTLS(config.BindAddr, certFile, keyFile, srv)
}

func messageSender(logger *logrus.Logger) {
	conn, ch := bus.GetChannel(addr)
	defer conn.Close()
	defer ch.Close()

	for {
		msg := "Golang is awesome - Keep Moving Forward!"
		bus.SendTextMessage(ch, queueName, msg)
		// logger.Printf(" [x] Congrats, sending message: %s", msg)
	}
}

func messageReceiver(logger *logrus.Logger) {
	conn, ch := bus.GetChannel(addr)
	defer conn.Close()
	defer ch.Close()

	msgs := bus.ReceiveMessage(ch, queueName)
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
