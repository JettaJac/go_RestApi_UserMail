package store_test

import (
	"main/internal/model"
	"main/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@mail.com",
	})
	// fmt.Println(u, err)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@mail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@mail.com",
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
