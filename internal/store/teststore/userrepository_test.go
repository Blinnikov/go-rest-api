package teststore_test

import (
	"testing"

	"github.com/blinnikov/go-rest-api/internal/app/model"
	"github.com/blinnikov/go-rest-api/internal/store"
	"github.com/blinnikov/go-rest-api/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_ReturnsErrorForNoUser(t *testing.T) {
	s := teststore.New()

	email := "Toto.Cutugno@sanremo.it"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}

func TestUserRepository_Find_ReturnsErrorForNoUser(t *testing.T) {
	s := teststore.New()

	id := 100500
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}

func TestUserRepository_FindByEmail_ReturnsUser(t *testing.T) {
	s := teststore.New()

	email := "Toto.Cutugno@sanremo.it"
	tu := model.TestUser(t)
	tu.Email = email
	s.User().Create(tu)

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, email, u.Email)
}

func TestUserRepository_Find_ReturnsUser(t *testing.T) {
	s := teststore.New()

	tu := model.TestUser(t)
	s.User().Create(tu)
	id := tu.ID

	u, err := s.User().Find(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, id, u.ID)
}
