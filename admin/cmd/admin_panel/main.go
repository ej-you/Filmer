package main

import (
	"log"

	"Filmer/admin/config"
	"Filmer/admin/internal/app"
)

func main() {
	if err := startApp(); err != nil {
		log.Fatal(err)
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
