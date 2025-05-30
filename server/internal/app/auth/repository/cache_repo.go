package repository

import (
	"net/http"

	"Filmer/server/config"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/httperror"

	"Filmer/server/internal/app/auth"
)

const _blacklistKeyPrefix = "token:blacklisted:" // key prefix for blacklisted tokens values

var _blacklistValue = []byte("t") // value for blacklisted tokens

var _ auth.CacheRepo = (*authCacheRepository)(nil)

// auth.CacheRepo implementation.
type authCacheRepository struct {
	cfg          *config.Config
	cacheStorage cache.Storage
}

// Returns auth.CacheRepo interface.
func NewCacheRepository(cfg *config.Config, cacheStorage cache.Storage) auth.CacheRepo {
	return &authCacheRepository{
		cfg:          cfg,
		cacheStorage: cacheStorage,
	}
}

// Set token to blacklist expiring at cfg.App.TokenExpired time.
func (r authCacheRepository) SetTokenToBlacklist(token string) error {
	err := r.cacheStorage.Set(_blacklistKeyPrefix+token, _blacklistValue, r.cfg.App.TokenExpired)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to add token to blacklist", err)
	}
	return nil
}

// Get bool value by given key-token.
// Returns true, if given token is blacklisted.
func (r authCacheRepository) TokenIsBlacklisted(token string) (bool, error) {
	blacklistedValue, err := r.cacheStorage.Get(_blacklistKeyPrefix + token)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to get blacklisted token", err)
	}
	return len(blacklistedValue) > 0, nil
}
