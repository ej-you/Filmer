package http

import (
	"errors"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/internal/pkg/logger"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.NewLogger().Error(err)

	var statusCode, message string

	var fiberErr *fiber.Error
	switch {
	// if *fiber.Error error
	case errors.As(err, &fiberErr):
		statusCode = strconv.Itoa(fiberErr.Code)
		message = fiberErr.Message
	default:
		statusCode = "500"
		message = err.Error()
	}

	// render 404 page for NotFound error
	if statusCode == "404" {
		return ctx.Status(404).Render("404", fiber.Map{})
	}

	// url and query params before error occurs
	url := ctx.OriginalURL()
	queryParams := ctx.Queries()
	// add query params for redirect bask with error message
	queryParams["statusCode"] = statusCode
	queryParams["message"] = message

	// redirect back with error message query-params
	return ctx.RedirectToRoute(url, fiber.Map{"queries": queryParams}, 303)
}
