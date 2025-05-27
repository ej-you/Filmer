// Server binary starts server.
package main

import (
	"log"

	"Filmer/server/config"
	"Filmer/server/internal/app/server"
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
	// init server
	server := server.New(cfg)
	// if err != nil {
	// 	return err
	// }
	// run server
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
