package teststore_test

import (
	"main/internal/model"
	"main/internal/store"
	"main/internal/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err) //err
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New() // new store
	email := "user@mail.com"
	_, err := s.User().FindByEmail(email)
	// assert.Error(t, err)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error()) //err

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
