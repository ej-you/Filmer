package utils

import (
	"fmt"
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
func ParseUserIDFromContext(ctx *fiber.Ctx) (uuid.UUID, error) {
	// parse token from ctx
	accessToken, ok := ctx.Locals("accessToken").(*jwt.Token)
	if !ok {
		return uuid.Nil, httpError.NewHTTPError(500, "failed to parse access token", fmt.Errorf("failed to parse access token"))
	}
	// parse user ID from token as string
	stringUserID, err := accessToken.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, httpError.NewHTTPError(500, "failed to parse user id from token", err)
	}
	// convert string user ID to UUID
	uuidUserID, err := uuid.Parse(stringUserID)
	if err != nil {
		return uuid.Nil, httpError.NewHTTPError(500, "failed to parse user id", err)
	}
	return uuidUserID, nil
}

// Parse token saved in context
func ParseRawTokenFromContext(ctx *fiber.Ctx) string {
	return ctx.Locals("accessToken").(*jwt.Token).Raw
}
