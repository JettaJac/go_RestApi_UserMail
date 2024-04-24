package app

// Config is the application configuration
type Config struct {
	BindAddr string `toml:"bind_addr"` // The bind address to listen on
	LogLevel string `toml:"log_level"` // The log level to use
}

// NewConfig creates a new configuration
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
