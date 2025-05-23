package database

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"Filmer/server/config"
	"Filmer/server/internal/pkg/logger"
)

var once sync.Once

// DB connection
var dbConnection *gorm.DB

// DB connection constructor
func NewCockroachClient(cfg *config.Config, appLogger logger.Logger) *gorm.DB {
	var err error

	once.Do(func() {
		appLogger.Infof("Process %d is connecting to CockroachDB node with %q", os.Getpid(), cfg.Database.ConnString)

		dbConnection, err = gorm.Open(postgres.New(postgres.Config{
			DSN: cfg.Database.ConnString,
		}), &gorm.Config{
			// set UTC time zone
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			// disable NotFound errors logging
			Logger: gormlogger.New(
				log.New(cfg.LogOutput.Error, "[SQL ERROR]\t", log.Ldate|log.Ltime),
				gormlogger.Config{
					LogLevel:                  gormlogger.Warn,
					IgnoreRecordNotFoundError: true,
					Colorful:                  false,
				},
			),
		})

		if err != nil {
			panic(err)
		}
		appLogger.Infof("Process %d successfully connected to CockroachDB!", os.Getpid())
	})
	return dbConnection
}
