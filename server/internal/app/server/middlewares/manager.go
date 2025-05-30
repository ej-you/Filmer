// Package middleware provides middleware manager with all server middlewares.
package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	fiberjwt "github.com/gofiber/contrib/jwt"
	fiberswagger "github.com/gofiber/contrib/swagger"
	fiber "github.com/gofiber/fiber/v2"
	fibercache "github.com/gofiber/fiber/v2/middleware/cache"
	fibercors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/auth"
	authRepository "Filmer/server/internal/app/auth/repository"
	authUsecase "Filmer/server/internal/app/auth/usecase"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/errhandler"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/token"
	"Filmer/server/internal/pkg/utils"
)

var _ MiddlewareManager = (*middlewareManager)(nil)

// Interface with all necessary middlewares for server.
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CORS() fiber.Handler
	Cache() fiber.Handler
	Swagger() fiber.Handler
	JWTAuth() fiber.Handler
}

// MiddlewareManager implementation.
type middlewareManager struct {
	cfg          *config.Config
	cacheStorage cache.Storage
	authUC       auth.Usecase
}

// MiddlewareManager constructor.
func NewMiddlewareManager(cfg *config.Config,
	dbClient *gorm.DB, cacheStorage cache.Storage) MiddlewareManager {

	authRepo := authRepository.NewDBRepo(dbClient)
	authCacheRepo := authRepository.NewCacheRepository(cfg, cacheStorage)
	authUsecase := authUsecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &middlewareManager{
		cfg:          cfg,
		cacheStorage: cacheStorage,
		authUC:       authUsecase,
	}
}

// Middleware for logging requests to server.
func (m middlewareManager) Logger() fiber.Handler {
	logFormat := "${time} | ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n"

	return fiberlogger.New(fiberlogger.Config{
		TimeFormat:    "2006-01-02T15:04:05-0700",
		Format:        logFormat,
		Output:        m.cfg.LogOutput.Info,
		DisableColors: false,
	})
}

// Middleware for panic recovery for continuous work.
func (m middlewareManager) Recover() fiber.Handler {
	return fiberrecover.New()
}

// CORS middleware.
func (m middlewareManager) CORS() fiber.Handler {
	return fibercors.New(fibercors.Config{
		AllowOrigins: m.cfg.App.CorsAllowedOrigins,
		AllowMethods: m.cfg.App.CorsAllowedMethods,
	})
}

// Cache middleware.
// Used after JWTAuth middleware for caching "search movies" and "person info".
func (m middlewareManager) Cache() fiber.Handler {
	return fibercache.New(fibercache.Config{
		Storage: m.cacheStorage,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.OriginalURL()
		},
		ExpirationGenerator: func(_ *fiber.Ctx, _ *fibercache.Config) time.Duration {
			return utils.ToNextDayDuration(time.Now().UTC())
		},
	})
}

// Middleware for Swagger docs.
func (m middlewareManager) Swagger() fiber.Handler {
	return fiberswagger.New(fiberswagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    fmt.Sprintf("%s docs", m.cfg.App.Name),
	})
}

// Middleware for parsing access token from headers to context and validate it.
func (m middlewareManager) JWTAuth() fiber.Handler {
	return fiberjwt.New(fiberjwt.Config{
		ContextKey:     "accessToken",
		SigningKey:     fiberjwt.SigningKey{Key: m.cfg.App.JwtSecret},
		SuccessHandler: m.checkBlacklistedToken(),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			switch {
			// if token expired error
			case errors.Is(err, jwt.ErrTokenExpired):
				err = httperror.New(http.StatusForbidden,
					"token is expired", err)
			// if token is missing
			case errors.Is(err, fiberjwt.ErrJWTMissingOrMalformed):
				err = httperror.New(http.StatusUnauthorized,
					"token is missing or malformed", err)
			}
			return errhandler.CustomErrorHandler(ctx, err)
		},
	})
}

// Next-step middleware after JWTAuth for restrict user access with blacklisted tokens.
func (m middlewareManager) checkBlacklistedToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := token.ParseRawTokenFromContext(ctx)
		// check token is blacklisted
		if err := m.authUC.RestrictBlacklistedToken(token); err != nil {
			return err
		}
		return ctx.Next()
	}
}
