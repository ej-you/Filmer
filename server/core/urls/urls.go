package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskUrls "server/app_kinopoisk/urls"
)


func InitRoutes(router fiber.Router) {
	kinopoiskUrls.AppRoutes(router.Group("/kinopoisk/films"))
}
