// Migrator binary is a migration manager for server DB.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"

	"Filmer/server/cmd/migrator/commands"
	"Filmer/server/config"
	"Filmer/server/pkg/migrate"
)

func main() {
	if err := startMigrator(); err != nil {
		log.Fatal(err)
	}
}

func startMigrator() error {
	// load config
	cfg, err := config.New()
	if err != nil {
		return err
	}
	// create migrate manager
	migrateManager, err := migrate.NewCockroachMigrate(cfg.Database.MigrationsURL, cfg.Database.ConnURL)
	if err != nil {
		return err
	}
	// defer migrate manager close
	defer func() {
		if err := migrateManager.Close(); err != nil {
			log.Printf("close migrate manager: %v \n", err)
		}
	}()

	// create migrator cmd
	cmd := &cli.Command{
		Name:  "migrator",
		Usage: "Migration manager for application DB",
		Commands: []*cli.Command{
			commands.NewStatus(migrateManager),
			commands.NewDown(migrateManager),
			commands.NewUp(migrateManager),
			commands.NewForce(migrateManager),
		},
	}
	// run migrator cmd
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		return fmt.Errorf("migrator cmd: %w", err)
	}
	return nil
}
