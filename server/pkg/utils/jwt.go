package utils

import (
	"time"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"

	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/config"
)


// Generate new access token for user
func ObtainToken(cfg *config.Config, userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().UTC().Add(cfg.App.TokenExpired).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(cfg.App.JwtSecret))
	if err != nil {
		return "", httpError.NewHTTPError(500, "failed to obtain token", err)
	}
	return tokenString, nil
}


// Parse user ID from token saved in context
func ParseUserIDFromContext(ctx *fiber.Ctx) uuid.UUID {
	stringUserID, _ := ctx.Locals("accessToken").(*jwt.Token).Claims.GetSubject()
	// claims := accessToken.Claims.(jwt.MapClaims)
	// stringUserID := claims["userID"].(string)
	return uuid.MustParse(stringUserID)
}

// Parse token saved in context
func ParseRawTokenFromContext(ctx *fiber.Ctx) string {
	return ctx.Locals("accessToken").(*jwt.Token).Raw
}
