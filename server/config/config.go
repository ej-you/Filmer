// Package config provides loading config data from
// external sources like env variables.
package config

import (
	"fmt"
	"io"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type (
	Config struct {
		App
		Cache
		Database
		RabbitMQ
		KinopoiskAPI
		LogOutput
	}

	App struct {
		Name               string        `env:"APP_NAME" env-default:"Filmer API" env-description:"app name for server (default: Filmer API)"`
		Port               string        `env:"SERVER_PORT" env-default:"8000" env-description:"server port (default: 8000)"`
		LogDir             string        `env:"LOG_DIR" env-default:"." env-description:"directory for log files (default: .)"`
		CorsAllowedOrigins string        `env-required:"true" env:"SERVER_CORS_ALLOWED_ORIGINS" env-description:"cors allowed origins"`
		CorsAllowedMethods string        `env-required:"true" env:"SERVER_CORS_ALLOWED_METHODS" env-description:"cors allowed methods"`
		JwtSecret          []byte        `env-required:"true" env:"JWT_SECRET" env-description:"secret for JWT-token signature"`
		TokenExpired       time.Duration `env:"TOKEN_EXPIRED" env-default:"30m" env-description:"JWT-token expired duration (default: 30m)"`
		KeepAliveTimeout   time.Duration `env:"KEEP_ALIVE_TIMEOUT" env-default:"60s"`
	}

	Cache struct {
		Host       string `env-required:"true" env:"REDIS_HOST" env-description:"redis host"`
		Port       string `env-required:"true" env:"REDIS_PORT" env-description:"redis port"`
		ConnString string
	}

	Database struct {
		MigrationsURL string `env:"MIGRATIONS_URL" env-default:"file://migrations"`
		User          string `env-required:"true" env:"DB_USER" env-description:"db user"`
		Host          string `env-required:"true" env:"DB_HOST" env-description:"db host"`
		Port          string `env-required:"true" env:"DB_PORT" env-description:"db port"`
		Name          string `env-required:"true" env:"DB_NAME" env-description:"db name"`
		ConnString    string
		ConnURL       string
	}

	RabbitMQ struct {
		User     string `env-required:"true" env:"RABBITMQ_DEFAULT_USER" env-description:"RabbitMQ user"`
		Password string `env-required:"true" env:"RABBITMQ_DEFAULT_PASS" env-description:"RabbitMQ name"`
		Host     string `env-required:"true" env:"RABBITMQ_HOST" env-description:"RabbitMQ host"`
		Port     string `env-required:"true" env:"RABBITMQ_PORT" env-description:"RabbitMQ port"`
		ConnURL  string
	}

	KinopoiskAPI struct {
		UnofficialKey string        `env-required:"true" env:"KINOPOISK_API_UNOFFICIAL_KEY" env-description:"key from Kinopoisk API Unofficial"`
		Key           string        `env-required:"true" env:"KINOPOISK_API_KEY" env-description:"key from Kinopoisk API"`
		DataExpired   time.Duration `env:"KINOPOISK_API_DATA_EXPIRED" env-default:"360h" env-description:"kinopoisk API data update duration (default: 360h)"`
	}

	LogOutput struct {
		Info  io.Writer
		Error io.Writer
	}
)

// New returns app config loaded from ENV-vars.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}
	cfg.Cache.ConnString = fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port)
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
	cfg.RabbitMQ.ConnURL = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	cfg.LogOutput.Info = &lumberjack.Logger{
		Filename:   cfg.App.LogDir + "/info.log",
		MaxSize:    1, // megabytes
		MaxBackups: 10,
		Compress:   false,
	}
	// logging errors with log rotation
	cfg.LogOutput.Error = &lumberjack.Logger{
		Filename:   cfg.App.LogDir + "/error.log",
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		Compress:   false,
	}
	return cfg, nil
}
