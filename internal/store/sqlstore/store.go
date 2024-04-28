package sqlstore

import (
	"database/sql"
	"main/internal/store"

	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Store is the structure for the store
type Store struct {
	//config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New creates a new store“
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// // Open database connection
// func (s *Store) Open() error {
// 	db, err := sql.Open("postgres", s.config.DatabaseURL)
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return err
// 	}

// 	s.db = db
// 	// ____
// 	// m, err := migrate.New(
// 	// 	"github://mattes:personal-access-token@mattes/migrate_test",
// 	// 	"postgres://localhost:5432/database?sslmode=enable")
// 	// m.Steps(2)
// 	// ____
// 	return nil
// }

// // Close database connection
// func (s *Store) Close() {
// 	s.db.Close()
// }

// User returns a user repository for the store
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
