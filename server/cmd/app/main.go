package main

import (
	"os"

	"Filmer/server/cmd/migrate"
	"Filmer/server/config"
	"Filmer/server/internal/app/server"
)

func main() {
	args := os.Args
	// if "migrate" arg is presented
	if len(args) > 1 && args[1] == "migrate" {
		migrate.RunMigrates()
		return
	}

	cfg := config.NewConfig()
	server := server.NewServer(cfg)

	server.Run()
}
