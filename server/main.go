package main

import (
	"fmt"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ej-you/go-utils/env"

	coreErrors "server/core/errors"
	coreServices "server/core/services"
	coreUrls "server/core/urls"
	"server/db"
	"server/settings"
)


func main() {
	// проверка, что эти переменные окружения заданы
	env.MustBePresented(
		"SERVER_PORT", "JWT_SECRET",
		"REDIS_HOST", "REDIS_PORT",
		"DB_USER", "DB_HOST", "DB_PORT", "DB_NAME",
		"SERVER_CORS_ALLOWED_ORIGINS", "SERVER_CORS_ALLOWED_METHODS",
		"KINOPOISK_API_UNOFFICIAL_KEY", "KINOPOISK_API_KEY",
	)

	// если при запуске указан аргумент "migrate"
	args := os.Args
	if len(args) > 1 {
		// проведение миграций БД без запуска самого приложения
		if args[1] == "migrate" {
			db.Migrate()
			return
		}
	}

	// инициализация приложения
	app := fiber.New(fiber.Config{
		AppName: "Filmer API v1.0.0",
		ErrorHandler: coreErrors.CustomErrorHandler,
		JSONEncoder: coreServices.EasyjsonEncoder,
		JSONDecoder: coreServices.EasyjsonDecoder,
		// https://www.f5.com/company/blog/nginx/socket-sharding-nginx-release-1-9-1
		Prefork: true,
		ServerHeader:  "Filmer API",

	})

	// логгер
	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		Format: "${time} | ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	}))
	// восстановление паник
	app.Use(fiberRecover.New())

	// подключение CORS
	app.Use(fiberCORS.New(fiberCORS.Config{
	    AllowOrigins: settings.CorsAllowedOrigins,
	    AllowMethods: settings.CorsAllowedMethods,
	}))

	api := app.Group("/api/v1")
	// настройка URL
	coreUrls.InitRoutes(api)

	settings.ErrorLog.Fatal(app.Listen(fmt.Sprintf(":%s", settings.Port)))
}
