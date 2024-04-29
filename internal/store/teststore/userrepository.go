package teststore

import (
	// "errors"
	"main/internal/model"
	"main/internal/store"
)

// UserRepository test
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create a new user in the repository test
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		// return nil, errors.New("user not found")
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
