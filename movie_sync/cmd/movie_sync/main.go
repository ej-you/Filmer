package main

import (
	"Filmer/movie_sync/config"
	"Filmer/movie_sync/internal/app"
	"Filmer/movie_sync/internal/pkg/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	// init logger
	logger.Init()

	if err := startApp(); err != nil {
		logrus.Fatal(err)
	}
}

func startApp() error {
	// load config
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// init application
	application, err := app.New(cfg)
	if err != nil {
		return err
	}
	// run application
	if err := application.Run(); err != nil {
		return err
	}
	return nil
}
