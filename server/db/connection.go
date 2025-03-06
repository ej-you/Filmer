package db

import (
	"os"
	"sync"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	gormLogger "gorm.io/gorm/logger"

	"server/db/schemas"
	"server/settings"
)


// ошибка ненахождения записи
var NotFoundError = gorm.ErrRecordNotFound


var once sync.Once
// соединение с БД
var dbConnection *gorm.DB

// функция для получения соединения с БД
func GetConn() *gorm.DB {
	var err error
	
	once.Do(func() {
		settings.InfoLog.Printf("Process %d is connecting to CockroachDB node with %q", os.Getpid(), settings.CockroachNodeAddr)

		dbConnection, err = gorm.Open(postgres.New(postgres.Config{
			DSN: settings.CockroachNodeAddr,
		}), &gorm.Config{
			// выставляем временную зону в UTC
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			// отключаем логирование NotFound ошибок
			Logger: gormLogger.New(
				settings.ErrorLog,
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
		settings.InfoLog.Printf("Process %d successfully connected to CockroachDB!", os.Getpid())
	})
	return dbConnection
}


// создание таблиц в БД по структурам в Go
func Migrate() {
	var err error

	conn := GetConn()
	settings.InfoLog.Println("Start migration...")

	settings.InfoLog.Println("Migrate \"User\" schema...")
	err = conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic(err)
	}

	settings.InfoLog.Println("Migrate \"Movie\" schema...")
	err = conn.AutoMigrate(&schemas.Movie{})
	if err != nil {
		panic(err)
	}

	settings.InfoLog.Println("Migrate \"Genre\" schema...")
	err = conn.AutoMigrate(&schemas.Genre{})
	if err != nil {
		panic(err)
	}

	settings.InfoLog.Println("Migrate \"UserMovie\" schema...")
	err = conn.AutoMigrate(&schemas.UserMovie{})
	if err != nil {
		panic(err)
	}

	settings.InfoLog.Println("CockroachDB -- Migrated successfully!")
}
