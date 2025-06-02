package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		RestAPI
	}

	App struct {
		Name         string        `env:"APP_NAME" env-default:"Filmer admin-panel" env-description:"app name for client (default: Filmer admin-panel)"`
		Port         string        `env:"ADMIN_PANEL_PORT" env-default:"8080" env-description:"admin-panel port (default: 8080)"`
		TokenExpired time.Duration `env:"TOKEN_EXPIRED" env-default:"30m" env-description:"REST API token expired duration (default: 30m)"`
		CookieSecure bool          `env:"COOKIES_SECURE" env-default:"false" env-description:"Set secure=true for cookies (default: false)"`
	}

	RestAPI struct {
		Host string `env-required:"true" env:"REST_API_HOST" env-description:"host addr for REST API"`
	}
)

// New returns app config loaded from ENV-vars.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load env-variables: %w", err)
	}
	return cfg, nil
}
