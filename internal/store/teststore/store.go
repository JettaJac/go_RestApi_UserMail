package teststore

import (
	"main/internal/model"
	"main/internal/store"

	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Store is the structure for the store
type Store struct {
	userRepository *UserRepository
}

// New creates a new store“
func New() *Store {
	return &Store{}
}

// User returns a user repository for the store
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
