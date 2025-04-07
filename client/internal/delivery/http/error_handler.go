package http

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/internal/pkg/logger"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.NewLogger().Error(err)

	var fiberErr *fiber.Error
	// if unknown error
	if !errors.As(err, &fiberErr) {
		return ctx.Status(http.StatusInternalServerError).Render("500", fiber.Map{"message": err.Error()})
	}

	// render 404 page for NotFound error (exclude 404 from REST API)
	if strings.HasPrefix(fiberErr.Message, "Cannot GET") {
		return ctx.Status(http.StatusNotFound).Render("404", fiber.Map{})
	}
	// render 402 page for API limit error
	if fiberErr.Code == http.StatusPaymentRequired {
		return ctx.Status(http.StatusPaymentRequired).Render("402", fiber.Map{})
	}

	// url and query params before error occurs
	url := ctx.OriginalURL()
	queryParams := ctx.Queries()
	// add query params for redirect bask with error message
	queryParams["statusCode"] = strconv.Itoa(fiberErr.Code)
	queryParams["message"] = fiberErr.Message

	// redirect back with error message query-params
	return ctx.RedirectToRoute(url, fiber.Map{"queries": queryParams}, http.StatusSeeOther)
}
