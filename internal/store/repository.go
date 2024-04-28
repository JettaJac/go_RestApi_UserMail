package store

import "main/internal/model"

// UserRepository...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
