package sqlstore

import (
	"database/sql"

	"github.com/blinnikov/go-rest-api/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) WriteTime(time string) {
	s.db.Exec("INSERT INTO time (time) VALUES ($1)", time)
}
