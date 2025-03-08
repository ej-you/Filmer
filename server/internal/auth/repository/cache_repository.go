package repository

import (
	"Filmer/server/pkg/cache"
	"Filmer/server/config"
	httpError "Filmer/server/pkg/http_error"

	"Filmer/server/internal/auth"
)


const blacklistKeyPrefix = "token:blacklisted:" // key prefix for blacklisted tokens values


// auth.CacheRepository interface implementation
type authCacheRepository struct {
	cfg			*config.Config
	cacheClient	cache.Cache
}

// auth.CacheRepository constructor
// Returns auth.CacheRepository interface
func NewCacheRepository(cfg *config.Config, cacheClient cache.Cache) auth.CacheRepository {
	return &authCacheRepository{
		cfg: cfg,
		cacheClient: cacheClient,
	}
}

// Set token to blacklist expiring at cfg.App.TokenExpired time
func (this authCacheRepository) SetTokenToBlacklist(token string) error {
	err := this.cacheClient.Set(blacklistKeyPrefix+token, "true", this.cfg.App.TokenExpired)
	if err != nil {
	    return httpError.NewHTTPError(500, "failed to add token to blacklist", err)
	}
	return nil
}

// Get bool value by given key-token
// Returns true, if given token is blacklisted
func (this authCacheRepository) TokenIsBlacklisted(token string) (bool, error) {
	isBlacklisted, err := this.cacheClient.GetBool(blacklistKeyPrefix+token)
	if err != nil {
	    return false, httpError.NewHTTPError(500, "failed to get blacklisted token", err)
	}
	return isBlacklisted, nil
}
