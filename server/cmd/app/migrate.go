package main

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"Filmer/server/config"
)

func main() {
	cfg := config.NewConfig()

	// load migrations and connect to DB
	migrator, err := migrate.New("file://migrations", cfg.Database.ConnURL)
	if err != nil {
		panic(err)
	}
	// make all up migrations
	if err := migrator.Up(); err != nil {
		fmt.Println("ERROR:", err)
	}
	// close file and DB connection
	fileCloseErr, dbCloseErr := migrator.Close()
	if fileCloseErr != nil {
		panic(fileCloseErr)
	}
	if dbCloseErr != nil {
		panic(dbCloseErr)
	}
}
