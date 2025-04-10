package migrate

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb" // cockroachdb engine for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"          // file engine for migrate

	"Filmer/server/config"
)

func RunMigrates() {
	cfg := config.NewConfig()

	fmt.Printf("Migrate DB %q\n", cfg.Database.ConnURL)
	// load migrations and connect to DB
	migrator, err := migrate.New("file://migrations", cfg.Database.ConnURL)
	if err != nil {
		panic(err)
	}
	// make all up migrations
	if err := migrator.Up(); err != nil {
		fmt.Println("Warning: migrate error was handled:", err)
	}
	// close file and DB connection
	fileCloseErr, dbCloseErr := migrator.Close()
	if fileCloseErr != nil {
		panic(fileCloseErr)
	}
	if dbCloseErr != nil {
		panic(dbCloseErr)
	}
	fmt.Println("Migration is finished!")
}
