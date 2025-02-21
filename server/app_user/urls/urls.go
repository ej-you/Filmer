package urls

import (
	fiber "github.com/gofiber/fiber/v2"

	"server/app_user/handlers"
	"server/core/services"
)


func AppRoutes(router fiber.Router) {
	router.Post("/sign-up", handlers.SignUp)
	router.Post("/login", handlers.Login)

	router.Use(services.AccessTokenMiddleware)
	router.Use(services.BlacklistedTokenMiddleware)
	router.Post("/logout", handlers.Logout)
}
