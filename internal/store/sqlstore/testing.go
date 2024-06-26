package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

// TestDB runs the test suite against a database.
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	// config := NewConfig()
	// config.DatabaseURL = databaseURL
	// s := New(config)
	// if err := s.Open(); err != nil {
	// 	t.Fatal(err)
	// }
	// return s, func(tables ...string) {
	// 	if len(tables) > 0 {
	// 		if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
	// 			t.Fatal(err)
	// 		}
	// 	}
	// 	s.Close()
	// }
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}
		db.Close()
	}
}
