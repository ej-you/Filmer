// Package middleware provides middleware manager with all server middlewares.
package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	fiberJWT "github.com/gofiber/contrib/jwt"
	fiberSwagger "github.com/gofiber/contrib/swagger"
	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/auth"
	authRepository "Filmer/server/internal/app/auth/repository"
	authUsecase "Filmer/server/internal/app/auth/usecase"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/utils"
)

var _ MiddlewareManager = (*middlewareManager)(nil)

// Interface with all necessary middlewares for server.
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CORS() fiber.Handler
	Swagger() fiber.Handler
	JWTAuth() fiber.Handler
}

// MiddlewareManager implementation.
type middlewareManager struct {
	cfg    *config.Config
	authUC auth.Usecase
}

// MiddlewareManager constructor.
func NewMiddlewareManager(cfg *config.Config,
	dbClient *gorm.DB, cache cache.Cache) MiddlewareManager {

	authRepo := authRepository.NewDBRepo(dbClient)
	authCacheRepo := authRepository.NewCacheRepository(cfg, cache)
	authUsecase := authUsecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &middlewareManager{
		cfg:    cfg,
		authUC: authUsecase,
	}
}

// Middleware for logging requests to server.
func (m middlewareManager) Logger() fiber.Handler {
	logFormat := "${time} | ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n"

	return fiberLogger.New(fiberLogger.Config{
		TimeFormat:    "2006-01-02T15:04:05-0700",
		Format:        logFormat,
		Output:        m.cfg.LogOutput.Info,
		DisableColors: false,
	})
}

// Middleware for panic recovery for continuous work.
func (m middlewareManager) Recover() fiber.Handler {
	return fiberRecover.New()
}

// CORS middleware.
func (m middlewareManager) CORS() fiber.Handler {
	return fiberCORS.New(fiberCORS.Config{
		AllowOrigins: m.cfg.App.CorsAllowedOrigins,
		AllowMethods: m.cfg.App.CorsAllowedMethods,
	})
}

// Middleware for Swagger docs.
func (m middlewareManager) Swagger() fiber.Handler {
	return fiberSwagger.New(fiberSwagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    fmt.Sprintf("%s docs", m.cfg.App.Name),
	})
}

// Middleware for parsing access token from headers to context and validate it.
func (m middlewareManager) JWTAuth() fiber.Handler {
	return fiberJWT.New(fiberJWT.Config{
		ContextKey:     "accessToken",
		SigningKey:     fiberJWT.SigningKey{Key: []byte(m.cfg.App.JwtSecret)},
		SuccessHandler: m.checkBlacklistedToken(),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			switch {
			// if token expired error
			case errors.Is(err, jwt.ErrTokenExpired):
				err = httperror.NewHTTPError(http.StatusForbidden,
					"token is expired", err)
			// if token is missing
			case errors.Is(err, fiberJWT.ErrJWTMissingOrMalformed):
				err = httperror.NewHTTPError(http.StatusUnauthorized,
					"token is missing or malformed", err)
			}
			return utils.CustomErrorHandler(ctx, err)
		},
	})
}

// Next-step middleware after JWTAuth for restrict user access with blacklisted tokens.
func (m middlewareManager) checkBlacklistedToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := utils.ParseRawTokenFromContext(ctx)
		// check token is blacklisted
		if err := m.authUC.RestrictBlacklistedToken(token); err != nil {
			return err
		}
		return ctx.Next()
	}
}
