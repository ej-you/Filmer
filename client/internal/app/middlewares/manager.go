package middlewares

import (
	// "errors"
	// "fmt"
	// "net/http"

	// fiberJWT "github.com/gofiber/contrib/jwt"
	fiber "github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	// jwt "github.com/golang-jwt/jwt/v5"

	// "Filmer/client/internal/auth"
	// authRepository "Filmer/client/internal/auth/repository"
	// authUsecase "Filmer/client/internal/auth/usecase"

	"Filmer/client/config"
	// "Filmer/client/pkg/cache"
	// httpError "Filmer/client/pkg/http_error"
	// "Filmer/client/pkg/utils"
)

// Interface with all necessary middlewares for client
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CookieParser() fiber.Handler
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
		accessToken := ctx.Cookies("auth")
		email := ctx.Cookies("email")

		// set vars to context
		ctx.Locals("accessToken", accessToken)
		ctx.Locals("email", email)
		return ctx.Next()
	}
}

// Redirect to login if cookies is not specified
func (mm appMiddlewareManager) ToLoginIfNoCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies("auth")

		// if cookies is not specified
		if accessToken == "" {
			return ctx.Redirect("/user/login", 303)
		}
		return ctx.Next()
	}
}

// Redirect to profile if cookies is specified
func (mm appMiddlewareManager) ToProfileIfCookie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse cookies from request
		accessToken := ctx.Cookies("auth")

		// if cookies is specified
		if accessToken != "" {
			return ctx.Redirect("/user/profile", 303)
		}
		return ctx.Next()
	}
}
