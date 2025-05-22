package utils

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"

	httpError "Filmer/server/internal/pkg/http_error"
)

//easyjson:json
type errorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

// custom error handler for server
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var errResp errorResponse

	var fiberErr *fiber.Error
	var httpErr httpError.HTTPError

	switch {
	// if *fiber.Error error
	case errors.As(err, &fiberErr):
		errResp.StatusCode = fiberErr.Code
		errResp.Message = fiberErr.Message
	// if httpError.HTTPError error
	case errors.As(err, &httpErr):
		errResp.StatusCode = httpErr.StatusCode()
		errResp.Message = httpErr.UserFriendlyMessage()
	default:
		errResp.StatusCode = 500
		errResp.Message = err.Error()
	}
	// send error response
	return ctx.Status(errResp.StatusCode).JSON(errResp)
}
