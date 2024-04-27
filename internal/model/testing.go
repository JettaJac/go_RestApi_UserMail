package model

import "testing"

// TestUser is a test user for testing
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@test.com",
		Password: "password",
	}
}
