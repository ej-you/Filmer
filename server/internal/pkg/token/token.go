// Package token contains functions to obtain and parse JWT-tokens.
package token

import (
	"fmt"
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"Filmer/server/config"
	"Filmer/server/internal/pkg/httperror"
)

const _accessTokenCtxKey = "accessToken" // key for access token value in fiber context

// NewToken generates new access token for user.
func New(cfg *config.Config, userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().UTC().Add(cfg.App.TokenExpired).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(cfg.App.JwtSecret))
	if err != nil {
		return "", httperror.New(http.StatusInternalServerError,
			"failed to obtain token", err)
	}
	return tokenString, nil
}

// ParseUserIDFromContext parses user ID from token saved in context.
func ParseUserIDFromContext(ctx *fiber.Ctx) (uuid.UUID, error) {
	// parse token from ctx
	accessToken, ok := ctx.Locals(_accessTokenCtxKey).(*jwt.Token)
	if !ok {
		return uuid.Nil, httperror.New(http.StatusInternalServerError,
			"failed to parse access token", fmt.Errorf("failed to parse access token"))
	}
	// parse user ID from token as string
	stringUserID, err := accessToken.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, httperror.New(http.StatusInternalServerError,
			"failed to parse user id from token", err)
	}
	// convert string user ID to UUID
	uuidUserID, err := uuid.Parse(stringUserID)
	if err != nil {
		return uuid.Nil, httperror.New(http.StatusInternalServerError,
			"failed to parse user id", err)
	}
	return uuidUserID, nil
}

// ParseRawTokenFromContext parses token saved in context
func ParseRawTokenFromContext(ctx *fiber.Ctx) string {
	return ctx.Locals(_accessTokenCtxKey).(*jwt.Token).Raw
}
