package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	"server/app_user/handlers"
)


func AppRoutes(router fiber.Router) {
	router.Post("/sign-up", handlers.SignUp)
	router.Post("/login", handlers.Login)
}
