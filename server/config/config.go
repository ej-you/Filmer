package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		Cache
		Database
		KinopoiskAPI
	}

	App struct {
		Name               string        `env:"APP_NAME" env-default:"Filmer API" env-description:"app name for server (default: Filmer API)"`
		Port               string        `env:"SERVER_PORT" env-default:"8000" env-description:"server port (default: 8000)"`
		CorsAllowedOrigins string        `env-required:"true" env:"SERVER_CORS_ALLOWED_ORIGINS" env-description:"cors allowed origins"`
		CorsAllowedMethods string        `env-required:"true" env:"SERVER_CORS_ALLOWED_METHODS" env-description:"cors allowed methods"`
		JwtSecret          string        `env-required:"true" env:"JWT_SECRET" env-description:"secret for JWT-token signature"`
		TokenExpired       time.Duration `env:"TOKEN_EXPIRED" env-default:"30m" env-description:"JWT-token expired duration (default: 30m)"`
	}

	Cache struct {
		Host       string `env-required:"true" env:"REDIS_HOST" env-description:"redis host"`
		Port       string `env-required:"true" env:"REDIS_PORT" env-description:"redis port"`
		ConnString string
	}

	Database struct {
		User       string `env-required:"true" env:"DB_USER" env-description:"db user"`
		Host       string `env-required:"true" env:"DB_HOST" env-description:"db host"`
		Port       string `env-required:"true" env:"DB_PORT" env-description:"db port"`
		Name       string `env-required:"true" env:"DB_NAME" env-description:"db name"`
		ConnString string
		ConnURL    string
	}

	KinopoiskAPI struct {
		UnofficialKey string        `env-required:"true" env:"KINOPOISK_API_UNOFFICIAL_KEY" env-description:"key from Kinopoisk API Unofficial"`
		Key           string        `env-required:"true" env:"KINOPOISK_API_KEY" env-description:"key from Kinopoisk API"`
		DataExpired   time.Duration `env:"KINOPOISK_API_DATA_EXPIRED" env-default:"360h" env-description:"kinopoisk API data update duration (default: 360h)"`
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
		cfg.Database.ConnString = fmt.Sprintf(
			"user=%s host=%s port=%s dbname=%s sslmode=disable",
			cfg.Database.User,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		)
		cfg.Database.ConnURL = fmt.Sprintf(
			"cockroach://%s@%s:%s/%s?sslmode=disable",
			cfg.Database.User,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		)
		cfg.Cache.ConnString = fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port)
	})
	return cfg
}
