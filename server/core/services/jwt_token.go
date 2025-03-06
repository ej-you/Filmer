package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"

	jwt "github.com/golang-jwt/jwt/v5"
	fiberJWT "github.com/gofiber/contrib/jwt"

	"server/redis"
	"server/core/errors"
	"server/settings"
)


// middleware для парсинга access JWT-токена из заголовка в контекст
var AccessTokenMiddleware fiber.Handler = fiberJWT.New(fiberJWT.Config{
	ContextKey: "accessToken",
	SigningKey: fiberJWT.SigningKey{Key: []byte(settings.JwtSecret)},
	ErrorHandler: errors.CustomErrorHandler,
})


// middleware для отказа в запросе с авторизацией, если токен в чёрном списке в кэше
func BlacklistedTokenMiddleware(ctx *fiber.Ctx) error {
	accessToken, ok := ctx.Locals("accessToken").(*jwt.Token)
	if !ok {
		return fiber.NewError(500, "failed to parse token")
	}
	
	// проверяем токен на нахождение в чёрном списке
	isBlacklisted, err := redis.GetBlacklistedToken(accessToken.Raw)
	if err != nil {
		return fmt.Errorf("blacklisted token middleware: %w", err)
	}
	// если токен в чёрном списке
	if isBlacklisted {
		return jwt.ErrTokenNotValidYet
	}
	return ctx.Next()
}


// выдача нового токена для юзера
func ObtainToken(userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp": time.Now().UTC().Add(settings.TokenExpiredTime).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(settings.JwtSecret))
	if err != nil {
		return "", fiber.NewError(500, "failed to obtain token: " + err.Error())
	}
	return tokenString, nil
}

// парсинг ID юзера из токена, сохранённого в контексте
func ParseUserIDFromContext(ctx *fiber.Ctx) uuid.UUID {
	accessToken := ctx.Locals("accessToken").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)

	stringUserID := claims["userID"].(string)
	return uuid.MustParse(stringUserID)
}

// парсинг токена, сохранённого в контексте
func ParseRawTokenFromContext(ctx *fiber.Ctx) string {
	return ctx.Locals("accessToken").(*jwt.Token).Raw
}
