package app

import (
	"bytes"
	"encoding/json"
	"main/internal/model"
	"main/internal/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/sessions"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
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
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
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
		{
			name: "innvalid payload",
			payloand: "invalid payload",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payloand: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payloand: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
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
