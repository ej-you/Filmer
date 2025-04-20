package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/repository"
)

// check given status code is 2xx
func StatusCode2xx(statusCode int) bool {
	return statusCode/100 == 2
}

// Create auth cookie for user
func GetAuthCookie(cfg *config.Config, authToken string) *fiber.Cookie {
	return createCookie("auth", authToken, cfg.App.PathPrefix, cfg.App.CookieSecure, cfg.App.TokenExpired)
}

// Clear auth cookie for user
func ClearAuthCookie(cfg *config.Config) *fiber.Cookie {
	return createCookie("auth", "", cfg.App.PathPrefix, cfg.App.CookieSecure, -time.Hour)
}

// Create cookie with user email
func GetEmailCookie(cfg *config.Config, email string) *fiber.Cookie {
	return createCookie("email", email, cfg.App.PathPrefix, cfg.App.CookieSecure, cfg.App.TokenExpired)
}

// Clear cookie with user email
func ClearEmailCookie(cfg *config.Config) *fiber.Cookie {
	return createCookie("email", "", cfg.App.PathPrefix, cfg.App.CookieSecure, -time.Hour)
}

func createCookie(name, value, path string, secure bool, expiresAfter time.Duration) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Strict",
		Expires:  time.Now().UTC().Add(expiresAfter),
	}
}

// Parse movie ID from path
func GetMovieIDPathParam(ctx *fiber.Ctx) (string, error) {
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return "", fiber.NewError(http.StatusBadRequest, "invalid movie ID was given")
	}
	return movieID, nil
}

// Get query params for category GET request
func GetCategoryQueryParams(ctx *fiber.Ctx) (repository.CategoryUserMoviesIn, error) {
	// parse query-params
	queryParams, err := url.ParseQuery(string(ctx.Request().URI().QueryString()))
	if err != nil {
		return nil, fmt.Errorf("parse query params: %w", err)
	}
	if len(queryParams["type"]) > 0 && queryParams["type"][0] == "все" {
		delete(queryParams, "type")
	}
	return repository.CategoryUserMoviesIn(queryParams), nil
}
