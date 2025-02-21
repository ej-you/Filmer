package settings

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)


// данные REST API
var Port string = os.Getenv("SERVER_PORT")
var JwtSecret string = os.Getenv("JWT_SECRET")

// время истечения действия токена
var TokenExpiredTime time.Duration = time.Minute * 5

// строка подключения к redis
var RedisAddr string = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

// строка подключения к БД
var CockroachNodeAddr string = fmt.Sprintf("user=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

// разрешённые источники и методы
var CorsAllowedOrigins string = os.Getenv("SERVER_CORS_ALLOWED_ORIGINS")
var CorsAllowedMethods string = os.Getenv("SERVER_CORS_ALLOWED_METHODS")

// Kinopoisk APIs
var KinopoiskApiUnofficialKey string = os.Getenv("KINOPOISK_API_UNOFFICIAL_KEY")
var KinopoiskApiKey string = os.Getenv("KINOPOISK_API_KEY")


// настройка логера ошибок
var ErrorLog *log.Logger = func() *log.Logger {
	// открытие файла для логирования ошибок
	file, err := os.OpenFile("./error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// настройка логирования в файл и в STDERR
	logWriter := io.MultiWriter(os.Stderr, file)
	return log.New(logWriter, "[ERROR]\t", log.Ldate|log.Ltime)
}()
// настройка инфо-логера
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
