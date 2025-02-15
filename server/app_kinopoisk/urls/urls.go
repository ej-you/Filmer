package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	"server/app_kinopoisk/handlers"
)


func AppRoutes(router fiber.Router) {
	router.Get("/search", handlers.SearchFilms)
	router.Get("/:id", handlers.GetFilmInfo)
	router.Get("/:id/staff", handlers.GetFilmStaff)
}
