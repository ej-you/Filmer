package db

import (
	"os"
	"sync"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"server/db/schemas"
	"server/settings"
)


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
		}), &gorm.Config{})
		
		if err != nil {
			panic(err)
		}
		settings.InfoLog.Printf("Process %d successfully connected to CockroachDB!", os.Getpid())
	})
	return dbConnection
}


// создание таблиц в БД по структурам в Go
func Migrate() {
	conn := GetConn()
	settings.InfoLog.Println("Start migration...")
	
	settings.InfoLog.Println("Migrate \"User\" schema...")
	err := conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic(err)
	}

	settings.InfoLog.Println("CockroachDB -- Migrated successfully!")
}
