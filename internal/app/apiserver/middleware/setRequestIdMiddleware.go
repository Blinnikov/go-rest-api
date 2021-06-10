package middleware

import (
	"context"
	"net/http"

	"github.com/blinnikov/go-rest-api/internal/app/apiserver/helpers"
	"github.com/google/uuid"
)

type SetRequestIdMiddleware struct {
	Next http.Handler
}

func (m *SetRequestIdMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	w.Header().Set("X-Request-ID", id)
	m.Next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), helpers.CtxKeyRequestID, id)))
}
