package store_test

import (
	"testing"

	"github.com/blinnikov/go-rest-api/internal/app/model"
	"github.com/blinnikov/go-rest-api/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")
	s := store.New(db)

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_ReturnsErrorForNoUser(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")
	s := store.New(db)

	email := "Toto.Cutugno@sanremo.it"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail_ReturnsUser(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")
	s := store.New(db)

	email := "Toto.Cutugno@sanremo.it"
	tu := model.TestUser(t)
	tu.Email = email
	s.User().Create(tu)

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, email, u.Email)
}
