package app

// import "main/internal/store"

// Config is the application configuration
type Config struct {
	BindAddr string `toml:"bind_addr"` // The bind address to listen on
	LogLevel string `toml:"log_level"` // The log level to use
	// Store    *store.Config // The store configuration
	DatabaseURL string `toml:"database_url"` // The database URL to use
	SessionKey string `toml:"session_key"` // The session key to use
}

// NewConfig creates a new configuration
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		// Store:    store.NewConfig(),
	}
}
