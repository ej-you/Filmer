package repository

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"Filmer/server/config"
	"Filmer/server/internal/app/user"
	"Filmer/server/internal/pkg/database"
)

var _dbRepo user.DBRepo

func TestMain(m *testing.M) {
	// load config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	// DB storage init
	dbStorage, err := database.New(cfg.Database.ConnString,
		database.WithIgnoreNotFound())
	if err != nil {
		log.Fatalf("connect to db: %v", err)
	}
	// create user repo
	_dbRepo = NewDBRepo(dbStorage)
	os.Exit(m.Run())
}

func TestGetActivity(t *testing.T) {
	t.Log("Get users activity")

	result, err := _dbRepo.GetActivity()
	require.NoError(t, err, "get activity error")

	t.Logf("Gotten result: %+v", result)
}
