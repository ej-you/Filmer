// Package config provides loading config data from
// external sources like env variables.
package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		RabbitMQ
		RestAPI
	}

	// App struct {
	// 	Name             string        `env:"APP_NAME" env-default:"Filmer admin-panel" env-description:"app name for client (default: Filmer admin-panel)"`
	// 	Port             string        `env:"ADMIN_PANEL_PORT" env-default:"8080" env-description:"admin-panel port (default: 8080)"`
	// 	KeepAliveTimeout time.Duration `env:"KEEP_ALIVE_TIMEOUT" env-default:"60s" env-description:"timeout for force shutdown (default: 60s)"`
	// }

	RabbitMQ struct {
		User     string `env-required:"true" env:"RABBITMQ_DEFAULT_USER" env-description:"RabbitMQ user"`
		Password string `env-required:"true" env:"RABBITMQ_DEFAULT_PASS" env-description:"RabbitMQ name"`
		Host     string `env-required:"true" env:"RABBITMQ_HOST" env-description:"RabbitMQ host"`
		Port     string `env-required:"true" env:"RABBITMQ_PORT" env-description:"RabbitMQ port"`
		ConnURL  string
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
	cfg.RabbitMQ.ConnURL = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	return cfg, nil
}
