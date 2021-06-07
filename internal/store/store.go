package store

type Store interface {
	User() UserRepository
	WriteTime(time string)
}
