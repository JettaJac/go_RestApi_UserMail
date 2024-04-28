package store

// User interface for user repository
type Store interface {
	User() UserRepository
}
