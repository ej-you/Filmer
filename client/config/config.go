package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		RestAPI
	}

	App struct {
		Name         string        `env:"APP_NAME" env-default:"Filmer client" env-description:"app name for client (default: Filmer client)"`
		Port         string        `env:"CLIENT_PORT" env-default:"7000" env-description:"client port (default: 7000)"`
		TokenExpired time.Duration `env:"TOKEN_EXPIRED" env-default:"30m" env-description:"REST API token expired duration (default: 30m)"`
		CookieSecure bool          `env:"COOKIES_SECURE" env-default:"false" env-description:"Set secure=true for cookies (default: false)"`
	}

	RestAPI struct {
		Host string `env-required:"true" env:"REST_API_HOST" env-description:"host addr for REST API"`
	}
)

var once sync.Once
var cfg = new(Config)

// Config constructor
// Returns app config loaded from ENV-vars
func NewConfig() *Config {
	once.Do(func() {
		if err := cleanenv.ReadEnv(cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
