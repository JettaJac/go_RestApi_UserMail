package store

import (
	"main/internal/model"
	// "main/internal/store"
)

// UserRepository is a repository for user
type UserRepository struct {
	store *Store
}

// Create a new user in the repository
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
