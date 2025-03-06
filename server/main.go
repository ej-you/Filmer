package main

import (
	"fmt"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ej-you/go-utils/env"

	coreErrors "server/core/errors"
	coreServices "server/core/services"
	coreUrls "server/core/urls"
	coreMiddlewares "server/core/middlewares"
	"server/db"
	"server/settings"
)

//	@title			Filmer API
//	@version		1.0.0
//	@description	This is a Filmer API for Kinopoisk API and DB

//	@license.name	MIT Licence
//	@license.url	https://github.com/ej-you/Filmer/blob/master/LICENCE

//	@host		127.0.0.1:3000
//	@basePath	/api/v1
//	@schemes	http

//	@accept						json
//	@produce					json
//	@query.collection.format	multi

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
//	@description				JWT security accessToken. Please, add it in the format "Bearer {AccessToken}" to authorize your requests.
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

	// настройка middlewares
	coreMiddlewares.SetupMiddlewares(app)
	// настройка URL
	coreUrls.InitRoutes(app)

	settings.ErrorLog.Fatal(app.Listen(fmt.Sprintf(":%s", settings.Port)))
}
