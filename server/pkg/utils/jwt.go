package utils

import (
	"time"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"

	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/config"
)


// выдача нового токена для юзера
func ObtainToken(cfg *config.Config, userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().UTC().Add(cfg.App.TokenExpired).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(cfg.App.JwtSecret))
	if err != nil {
		return "", httpError.NewHTTPError(500, "failed to obtain token: " + err.Error())
	}
	return tokenString, nil
}


// парсинг ID юзера из токена, сохранённого в контексте
func ParseUserIDFromContext(ctx *fiber.Ctx) uuid.UUID {
	stringUserID, _ := ctx.Locals("accessToken").(*jwt.Token).Claims.GetSubject()
	// claims := accessToken.Claims.(jwt.MapClaims)
	// stringUserID := claims["userID"].(string)
	return uuid.MustParse(stringUserID)
}

// парсинг токена, сохранённого в контексте
func ParseRawTokenFromContext(ctx *fiber.Ctx) string {
	return ctx.Locals("accessToken").(*jwt.Token).Raw
}
