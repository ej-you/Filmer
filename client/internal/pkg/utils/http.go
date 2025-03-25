package utils

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
)

func StatusCode2xx(statusCode int) bool {
	return statusCode/100 == 2
}

// create auth cookie for user
func GetAuthCookie(cfg *config.Config, authToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "auth",
		Value:    authToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   cfg.App.CookieSecure,
		SameSite: "Strict",
		Expires:  time.Now().Add(cfg.App.TokenExpired),
	}
}

// clear auth cookie for user
func ClearAuthCookie(cfg *config.Config) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   cfg.App.CookieSecure,
		SameSite: "Strict",
		Expires:  time.Now().Add(-time.Hour),
	}
}

// create cookie with user email
func GetEmailCookie(cfg *config.Config, email string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "email",
		Value:    email,
		Path:     "/",
		HTTPOnly: true,
		Secure:   cfg.App.CookieSecure,
		SameSite: "Strict",
		Expires:  time.Now().Add(cfg.App.TokenExpired),
	}
}

// clear cookie with user email
func ClearEmailCookie(cfg *config.Config) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "email",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   cfg.App.CookieSecure,
		SameSite: "Strict",
		Expires:  time.Now().Add(-time.Hour),
	}
}
