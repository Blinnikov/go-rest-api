package store_test

import (
	"testing"

	"github.com/blinnikov/go-rest-api/internal/app/model"
	"github.com/blinnikov/go-rest-api/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "igor.blinnikov@boost.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_ReturnsErrorForNoUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "Toto.Cutugno@sanremo.it"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail_ReturnsUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "Toto.Cutugno@sanremo.it"
	s.User().Create(&model.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, email, u.Email)
}
