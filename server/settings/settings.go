package settings

import (
	"io"
	"log"
	"os"
)


// данные REST API
var Port string = os.Getenv("SERVER_PORT")
var JwtSecret string = os.Getenv("JWT_SECRET")

// разрешённые источники и методы
var CorsAllowedOrigins string = os.Getenv("SERVER_CORS_ALLOWED_ORIGINS")
var CorsAllowedMethods string = os.Getenv("SERVER_CORS_ALLOWED_METHODS")

// Kinopoisk API
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
