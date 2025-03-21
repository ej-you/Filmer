package main

import (
	"Filmer/client/config"
	"Filmer/client/internal/app/client"
)

func main() {
	cfg := config.NewConfig()
	client := client.NewClient(cfg)

	client.Run()
}
