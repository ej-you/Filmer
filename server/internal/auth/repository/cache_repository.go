package repository

import (
	"fmt"

	"Filmer/server/pkg/cache"
	"Filmer/server/config"
	httpError "Filmer/server/pkg/http_error"

	"Filmer/server/internal/auth"
)


const blacklistKeyPrefix = "token:blacklisted:" // префикс ключа для значений черного списка


// тип интерфейса auth.CacheRepository
type authCacheRepository struct {
	cfg			*config.Config
	cacheClient	cache.Cache
}

// конструктор для типа интерфейса auth.CacheRepository
func NewCacheRepository(cfg *config.Config, cacheClient cache.Cache) auth.CacheRepository {
	return &authCacheRepository{
		cfg: cfg,
		cacheClient: cacheClient,
	}
}

// установка ключа-значения в кэш со временем просрочки как у токена
func (this authCacheRepository) SetTokenToBlacklist(token string) error {
	err := this.cacheClient.Set(blacklistKeyPrefix+token, "true", this.cfg.App.TokenExpired)
	if err != nil {
	    return httpError.NewHTTPError(500, "failed to add token to blacklist: " + err.Error())
	}
	return nil
}


// получение bool значения из кэша по ключу
func (this authCacheRepository) TokenIsBlacklisted(token string) (bool, error) {
	isBlacklisted, err := this.cacheClient.GetBool(blacklistKeyPrefix+token)
	if err != nil {
	    return false, fmt.Errorf("find in blacklist: %w", httpError.NewHTTPError(500, "failed to get token: " + err.Error()))
	}
	return isBlacklisted, nil
}
