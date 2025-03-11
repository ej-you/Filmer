package main

import (
	"Filmer/server/config"
	"Filmer/server/internal/app/server"
)

func main() {
	cfg := config.NewConfig()
	server := server.NewServer(cfg)

	server.Run()
}
