package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blinnikov/go-rest-api/internal/app/model"
	"github.com/blinnikov/go-rest-api/internal/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name         string
	payload      interface{}
	expectedCode int
}

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authnticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("test")))
	testCases := []TestCase{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@domain.com",
				"password": "s3cret",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "mess",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "email",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, runTestCase(s, tc, "/users"))
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("test")))

	testCases := []TestCase{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "mess",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "no-domain",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid passwrod",
			payload: map[string]string{
				"email":    u.Email,
				"password": "TryToHack",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, runTestCase(s, tc, "/sessions"))
	}
}

func runTestCase(s *server, tc TestCase, path string) func(*testing.T) {
	return func(t *testing.T) {
		rec := httptest.NewRecorder()
		b := &bytes.Buffer{}
		json.NewEncoder(b).Encode(tc.payload)
		req, _ := http.NewRequest(http.MethodPost, path, b)
		s.ServeHTTP(rec, req)
		assert.Equal(t, tc.expectedCode, rec.Code)
	}
}
