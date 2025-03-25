package http

import (
	fiber "github.com/gofiber/fiber/v2"
)

func indexGET(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}
