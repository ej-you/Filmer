package auth

import (
	fiber "github.com/gofiber/fiber/v2"
)


type Router interface {
	SetRoutes(router fiber.Router)
}
