package middlewares

import (
	"fmt"
	"errors"

	"gorm.io/gorm"
	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberJWT "github.com/gofiber/contrib/jwt"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/contrib/swagger"

	"Filmer/server/internal/auth"
	authUsecase "Filmer/server/internal/auth/usecase"
	authRepository "Filmer/server/internal/auth/repository"
	
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/utils"
	"Filmer/server/config"
)


// интерфейс со всеми необходимыми middleware для сервера
type MiddlewareManager interface {
	Logger() fiber.Handler
	Recover() fiber.Handler
	CORS() fiber.Handler
	Swagger() fiber.Handler
	JWTAuth() fiber.Handler
}


// хранит все необходимые middleware для сервера
type appMiddlewareManager struct {
	cfg		*config.Config
	authUC	auth.Usecase
}

// конструктор для типа интерфейса MiddlewareManager
func NewMiddlewareManager(cfg *config.Config, dbClient *gorm.DB, cache cache.Cache) MiddlewareManager {
	authRepo := authRepository.NewRepository(dbClient)
	authCacheRepo := authRepository.NewCacheRepository(cfg, cache)
	authUsecase := authUsecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &appMiddlewareManager{
		cfg: cfg,
		authUC: authUsecase,
	}
}

// middleware для логирования запросов к серверу
func (this appMiddlewareManager) Logger() fiber.Handler {
	return fiberLogger.New(fiberLogger.Config{
		TimeFormat: "2006-01-02T15:04:05-0700",
		Format: "${time} | pid ${pid} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
	})
}

// middleware для восстановления паник для непрерывной работы
func (this appMiddlewareManager) Recover() fiber.Handler {
	return fiberRecover.New()
}

// middleware для настройки CORS
func (this appMiddlewareManager) CORS() fiber.Handler {
	return fiberCORS.New(fiberCORS.Config{
		AllowOrigins: this.cfg.App.CorsAllowedOrigins,
		AllowMethods: this.cfg.App.CorsAllowedMethods,
	})
}

// middleware для Swagger документации
func (this appMiddlewareManager) Swagger() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath:	"/api/v1/",
		FilePath:	"./docs/swagger.json",
		Path:		"docs",
		Title:		fmt.Sprintf("%s docs", this.cfg.App.Name),
		CacheAge:	3600, // 1 hour
	})
}

// middleware для парсинга access JWT-токена из заголовка в контекст
func (this appMiddlewareManager) JWTAuth() fiber.Handler {
	return fiberJWT.New(fiberJWT.Config{
		ContextKey: "accessToken",
		SigningKey: fiberJWT.SigningKey{Key: []byte(this.cfg.App.JwtSecret)},
		SuccessHandler: this.checkBlacklistedToken(),
		ErrorHandler: func (ctx *fiber.Ctx, err error) error {
			switch {
				// если ошибка истёкшего токена
				case errors.Is(err, jwt.ErrTokenExpired):
					err = httpError.NewHTTPError(403, "token is expired")
				// если ошибка отсутствия токена
				case errors.Is(err, fiberJWT.ErrJWTMissingOrMalformed):
					err = httpError.NewHTTPError(401, "token is missing or malformed")
			}
			return utils.CustomErrorHandler(ctx, err)
		},
	})
}

// middleware для отказа в запросе с авторизацией, если токен в чёрном списке в кэше
func (this appMiddlewareManager) checkBlacklistedToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := utils.ParseRawTokenFromContext(ctx)
		// проверяем токен на нахождение в черном списке
		if err := this.authUC.RestrictBlacklistedToken(token); err != nil {
			return err
		}
		return ctx.Next()
	}
}
