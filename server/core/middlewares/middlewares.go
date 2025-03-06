package middlewares

import (
	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/contrib/swagger"

	"server/settings"
)


func SetupMiddlewares(app *fiber.App) {
	// логгер
	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		Format: "${time} | pid ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	}))
	// восстановление паник
	app.Use(fiberRecover.New())

	// подключение CORS
	app.Use(fiberCORS.New(fiberCORS.Config{
	    AllowOrigins: settings.CorsAllowedOrigins,
	    AllowMethods: settings.CorsAllowedMethods,
	}))

	// подключение Swagger документации
	app.Use(swagger.New(swagger.Config{
		BasePath:	"/api/v1/",
		FilePath:	"./docs/swagger.json",
		Path:		"docs",
		Title:		"Filmer API docs",
		CacheAge:	3600, // 1 hour
	}))
}
