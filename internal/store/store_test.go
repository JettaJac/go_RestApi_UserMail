package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv(("DATABASE_URL")) // set DATABASE_URL env var
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}
	os.Exit(m.Run())
}
