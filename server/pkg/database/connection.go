package database

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	gormLogger "gorm.io/gorm/logger"

	"Filmer/server/pkg/logger"
	"Filmer/server/config"
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
			Logger: gormLogger.New(
				log.New(os.Stderr, "[SQL ERROR]\t", log.Ldate|log.Ltime),
				gormLogger.Config{
					SlowThreshold: 200*time.Millisecond,
					LogLevel: gormLogger.Warn,
					IgnoreRecordNotFoundError: true,
					Colorful: true,
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
