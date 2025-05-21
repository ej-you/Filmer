package migrate

import (
	"errors"
	"fmt"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb" // cockroachdb engine for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"          // engine for migration files
)

var _ Migrate = (*cockroachMigrate)(nil)

// Migrate implementation.
type cockroachMigrate struct {
	mgrt *gomigrate.Migrate
}

func NewCockroachMigrate(sourceURL, databaseURL string) (Migrate, error) {
	mgrt, err := gomigrate.New(sourceURL, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create migrate manager: %w", err)
	}
	return &cockroachMigrate{mgrt: mgrt}, nil
}

func (c *cockroachMigrate) Status() (version uint, isDirty bool, err error) {
	v, d, err := c.mgrt.Version()
	if err != nil && !errors.Is(err, gomigrate.ErrNilVersion) {
		return 0, false, fmt.Errorf("migrate status: %w", err)
	}
	return v, d, nil
}

func (c *cockroachMigrate) Up() error {
	err := c.mgrt.Up()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func (c *cockroachMigrate) Down() error {
	err := c.mgrt.Down()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate down: %w", err)
	}
	return nil
}

func (c *cockroachMigrate) Step(n int) error {
	err := c.mgrt.Steps(n)
	if err != nil {
		return fmt.Errorf("migrate step: %w", err)
	}
	return nil
}

func (c *cockroachMigrate) Force(n int) error {
	err := c.mgrt.Force(n)
	if err != nil {
		return fmt.Errorf("migrate force: %w", err)
	}
	return nil
}

func (c *cockroachMigrate) Close() error {
	sourceCloseErr, dbCloseErr := c.mgrt.Close()
	if sourceCloseErr != nil && dbCloseErr != nil {
		return fmt.Errorf("close database: %s && close source: %w", dbCloseErr.Error(), sourceCloseErr)
	}
	if sourceCloseErr != nil {
		return fmt.Errorf("close source: %w", sourceCloseErr)
	}
	if dbCloseErr != nil {
		return fmt.Errorf("close database: %w", dbCloseErr)
	}
	return nil
}
