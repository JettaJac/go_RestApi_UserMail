package app

import (
	"database/sql"
	"main/internal/store/sqlstore"
	"net/http"
	// "github.com/golang-migrate/migrate/v4/database"
)

// Start app
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv.router)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

/*
import (
	"fmt"
	"io"
	"main/internal/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" force 20240426112011
// migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

// CREATE TABLE users (
//     id bigserial not null primary key,
//     email varchar not null unique,
//     encrypted_password varchar not null
// );
// DROP TABLE users;
// migrate create -ext sql -dir migrations create_users команда создающая миграции
// go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
//  postgres://jettajac:K.,jdm2018@host/restapi_dev
// postgres://jettajac:K.,jdm2018@host:port/restapi_dev?query
// migrate -source=postgres://jettajac:K.,jdm2018@localhost/restapi_dev -database=postgres://jettajac:K.,jdm2018@localhost/restapi_dev

// APIServer
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New app s
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Run app
func (s *APIServer) Run() error {
	fmt.Println("Start App")

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Start App")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
*/
