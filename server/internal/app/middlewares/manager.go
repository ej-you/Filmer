package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	fiberJWT "github.com/gofiber/contrib/jwt"
	fiberSwagger "github.com/gofiber/contrib/swagger"
	fiber "github.com/gofiber/fiber/v2"
	fiberCache "github.com/gofiber/fiber/v2/middleware/cache"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"Filmer/server/internal/auth"
	authRepository "Filmer/server/internal/auth/repository"
	authUsecase "Filmer/server/internal/auth/usecase"

	"Filmer/server/config"
	"Filmer/server/pkg/cache"
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/utils"
)

// Interface with all necessary middlewares for server
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CORS() fiber.Handler
	Cache() fiber.Handler
	Swagger() fiber.Handler
	JWTAuth() fiber.Handler
}

// MiddlewareManager interface implementation
type appMiddlewareManager struct {
	cfg    *config.Config
	authUC auth.Usecase
}

// MiddlewareManager constructor
func NewMiddlewareManager(cfg *config.Config, dbClient *gorm.DB, cache cache.Cache) MiddlewareManager {
	authRepo := authRepository.NewRepository(dbClient)
	authCacheRepo := authRepository.NewCacheRepository(cfg, cache)
	authUsecase := authUsecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &appMiddlewareManager{
		cfg:    cfg,
		authUC: authUsecase,
	}
}

// Middleware for logging requests to server
func (mm appMiddlewareManager) Logger() fiber.Handler {
	return fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		Format:     "${time} | pid ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	})
}

// Middleware for panic recovery for continuous work
func (mm appMiddlewareManager) Recover() fiber.Handler {
	return fiberRecover.New()
}

// CORS middleware
func (mm appMiddlewareManager) CORS() fiber.Handler {
	return fiberCORS.New(fiberCORS.Config{
		AllowOrigins: mm.cfg.App.CorsAllowedOrigins,
		AllowMethods: mm.cfg.App.CorsAllowedMethods,
	})
}

// Cache middleware
func (mm appMiddlewareManager) Cache() fiber.Handler {
	return fiberCache.New(fiberCache.Config{
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.OriginalURL()
		},
		Next: func(ctx *fiber.Ctx) bool {
			return ctx.Path() != "/api/v1/kinopoisk/films/search"
		},
		Expiration: mm.cfg.App.CacheExpiration,
	})
}

// Middleware for Swagger docs
func (mm appMiddlewareManager) Swagger() fiber.Handler {
	return fiberSwagger.New(fiberSwagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    fmt.Sprintf("%s docs", mm.cfg.App.Name),
	})
}

// Middleware for parsing access token from headers to context and validate it
func (mm appMiddlewareManager) JWTAuth() fiber.Handler {
	return fiberJWT.New(fiberJWT.Config{
		ContextKey:     "accessToken",
		SigningKey:     fiberJWT.SigningKey{Key: []byte(mm.cfg.App.JwtSecret)},
		SuccessHandler: mm.checkBlacklistedToken(),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			switch {
			// if token expired error
			case errors.Is(err, jwt.ErrTokenExpired):
				err = httpError.NewHTTPError(http.StatusForbidden, "token is expired", err)
			// if token is missing
			case errors.Is(err, fiberJWT.ErrJWTMissingOrMalformed):
				err = httpError.NewHTTPError(http.StatusUnauthorized, "token is missing or malformed", err)
			}
			return utils.CustomErrorHandler(ctx, err)
		},
	})
}

// Next-step middleware after JWTAuth for restrict user access with blacklisted tokens
func (mm appMiddlewareManager) checkBlacklistedToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := utils.ParseRawTokenFromContext(ctx)
		// check token is blacklisted
		if err := mm.authUC.RestrictBlacklistedToken(token); err != nil {
			return err
		}
		return ctx.Next()
	}
}
