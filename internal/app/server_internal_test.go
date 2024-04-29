package app

import (
	"bytes"
	"encoding/json"
	"main/internal/model"
	"main/internal/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCase := []struct {
		name         string
		payloand     interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payloand: map[string]string{
				"email":    "u@e1.ru",
				"password": "password",
			},
			expectedCode: http.StatusCreated, // 201, // StatusCreated
		},
		{
			name:         "invalid payload",
			payloand:     "invalid payload",
			expectedCode: 400, // StatusBadRequest
		},
		{
			name: "invalid params",
			payloand: map[string]string{
				"email":    "invalid",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payloand)
			req, _ := http.NewRequest("POST", "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store)
	testCase := []struct {
		name         string
		payloand     interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payloand: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payloand)
			req, _ := http.NewRequest("POST", "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
