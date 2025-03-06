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


// ошибка ненахождения записи
var NotFoundError = gorm.ErrRecordNotFound


var once sync.Once
// соединение с БД
var dbConnection *gorm.DB


// конструктор для получения соединения с БД
func NewCockroachClient(cfg *config.Config, appLogger logger.Logger) *gorm.DB {
	var err error
	
	once.Do(func() {
		appLogger.Infof("Process %d is connecting to CockroachDB node with %q", os.Getpid(), cfg.Database.ConnString)

		dbConnection, err = gorm.Open(postgres.New(postgres.Config{
			DSN: cfg.Database.ConnString,
		}), &gorm.Config{
			// выставляем временную зону в UTC
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			// отключаем логирование NotFound ошибок
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
