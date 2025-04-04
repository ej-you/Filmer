package repository

import (
	"net/http"

	"Filmer/server/config"
	"Filmer/server/pkg/cache"
	httpError "Filmer/server/pkg/http_error"

	"Filmer/server/internal/auth"
)

const blacklistKeyPrefix = "token:blacklisted:" // key prefix for blacklisted tokens values

// auth.CacheRepository interface implementation
type authCacheRepository struct {
	cfg         *config.Config
	cacheClient cache.Cache
}

// auth.CacheRepository constructor
// Returns auth.CacheRepository interface
func NewCacheRepository(cfg *config.Config, cacheClient cache.Cache) auth.CacheRepository {
	return &authCacheRepository{
		cfg:         cfg,
		cacheClient: cacheClient,
	}
}

// Set token to blacklist expiring at cfg.App.TokenExpired time
func (acr authCacheRepository) SetTokenToBlacklist(token string) error {
	err := acr.cacheClient.Set(blacklistKeyPrefix+token, "true", acr.cfg.App.TokenExpired)
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to add token to blacklist", err)
	}
	return nil
}

// Get bool value by given key-token
// Returns true, if given token is blacklisted
func (acr authCacheRepository) TokenIsBlacklisted(token string) (bool, error) {
	isBlacklisted, err := acr.cacheClient.GetBool(blacklistKeyPrefix + token)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to get blacklisted token", err)
	}
	return isBlacklisted, nil
}
