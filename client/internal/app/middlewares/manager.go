package middlewares

import (
	"fmt"
	"net/http"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"
	fiberCompress "github.com/gofiber/fiber/v2/middleware/compress"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	"Filmer/client/config"
	"Filmer/client/internal/app/constants"
)

// Interface with all necessary middlewares for client
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CookieParser() fiber.Handler
	Compression() fiber.Handler
	ToLoginIfNoCookie() fiber.Handler
	ToProfileIfCookie() fiber.Handler
}

// MiddlewareManager interface implementation
type appMiddlewareManager struct {
	cfg *config.Config
}

// MiddlewareManager constructor
func NewMiddlewareManager(cfg *config.Config) MiddlewareManager {
	return &appMiddlewareManager{
		cfg: cfg,
	}
}

// Middleware for logging requests to client
func (mm appMiddlewareManager) Logger() fiber.Handler {
	return fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	})
}

// Middleware for panic recovery for continuous work
func (mm appMiddlewareManager) Recover() fiber.Handler {
	return fiberRecover.New()
}

// Parsing access token and email from cookies to context (and redirect to login if cookies is not specified)
func (mm appMiddlewareManager) CookieParser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies(constants.CookieAuth)
		email := ctx.Cookies(constants.CookieEmail)

		// set vars to context
		ctx.Locals(constants.LocalsKeyAccessToken, accessToken)
		ctx.Locals(constants.LocalsKeyEmail, email)
		return ctx.Next()
	}
}

func (mm appMiddlewareManager) Compression() fiber.Handler {
	return fiberCompress.New(fiberCompress.Config{
		Level: fiberCompress.LevelBestSpeed,
	})
}

// Redirect to login if cookies is not specified
func (mm appMiddlewareManager) ToLoginIfNoCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies(constants.CookieAuth)

		// if cookies is not specified
		if accessToken == "" {
			// send next param (current url before redirect to login)
			redirectURL := fmt.Sprintf("/user/login?%s=%s", constants.NextQueryParam, url.QueryEscape(ctx.OriginalURL()))
			return ctx.Redirect(redirectURL, http.StatusSeeOther)
		}
		return ctx.Next()
	}
}

// Redirect to profile if cookies is specified
func (mm appMiddlewareManager) ToProfileIfCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies(constants.CookieAuth)

		// if cookies is specified
		if accessToken != "" {
			return ctx.Redirect("/user/profile", http.StatusSeeOther)
		}
		return ctx.Next()
	}
}
