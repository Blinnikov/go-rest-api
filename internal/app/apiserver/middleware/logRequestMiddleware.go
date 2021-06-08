package middleware

import (
	"net/http"
	"time"

	"github.com/blinnikov/go-rest-api/internal/app/apiserver/helpers"
	"github.com/sirupsen/logrus"
)

type LogRequestMiddleware struct {
	Next   http.Handler
	Logger *logrus.Logger
}

func (m *LogRequestMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := m.Logger.WithFields(logrus.Fields{
		"remote_addr": r.RemoteAddr,
		"request_id":  r.Context().Value(helpers.CtxKeyRequestID),
	})
	logger.Infof("Started %s %s", r.Method, r.RequestURI)

	start := time.Now()
	rw := &statusCodeResponseWriter{w, http.StatusOK}
	m.Next.ServeHTTP(rw, r)
	logger.Infof("Completed with %d %s in %v",
		rw.statusCode,
		http.StatusText(rw.statusCode),
		time.Since(start),
	)
}
