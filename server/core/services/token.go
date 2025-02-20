package services

import (
	"time"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"

	jwt "github.com/golang-jwt/jwt/v5"
	fiberJWT "github.com/gofiber/contrib/jwt"

	"server/core/errors"
	"server/settings"
)


// middleware для парсинга access JWT-токена из заголовка в контекст
var AccessTokenMiddleware fiber.Handler = fiberJWT.New(fiberJWT.Config{
	ContextKey: "accessToken",
	SigningKey: fiberJWT.SigningKey{Key: []byte(settings.JwtSecret)},
	ErrorHandler: errors.CustomErrorHandler,
})


// выдача новогог токена для юзера
func ObtainToken(userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp": time.Now().Add(settings.TokenExpiredTime).Unix(),
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
