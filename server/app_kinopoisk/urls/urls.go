package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	"server/app_kinopoisk/handlers"
	"server/core/services"
)


func AppRoutes(router fiber.Router) {
	router.Use(services.AccessTokenMiddleware)
	router.Use(services.BlacklistedTokenMiddleware)

	router.Get("/search", handlers.SearchFilms)
	router.Get("/:id", handlers.GetFilmInfo)
}
