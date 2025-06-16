// Package repository contains kinopoisk repository implementation.
package repository

import (
	"net/http"
	"time"

	"Filmer/server/internal/app/kinopoisk"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/utils"
)

const apiLimitPrefix = "api-limit-reached:" // key prefix for APIs limit values

var _apiLimitValue = []byte("t") // value for API limit

var _ kinopoisk.CacheRepo = (*cacheRepo)(nil)

// kinopoisk.CacheRepo implementation.
type cacheRepo struct {
	cacheStorage cache.Storage
}

// Returns kinopoisk.CacheRepo interface.
func NewCacheRepo(cacheStorage cache.Storage) kinopoisk.CacheRepo {
	return &cacheRepo{
		cacheStorage: cacheStorage,
	}
}

// Set api limit (for apiName API) expiring at the beginning of the next day (00:00 UTC).
func (c cacheRepo) SetAPILimit(apiName string) error {
	toNextDay := utils.ToNextDayDuration(time.Now().UTC())
	err := c.cacheStorage.Set(apiLimitPrefix+apiName, _apiLimitValue, toNextDay)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to set api limit", err)
	}
	return nil
}

// Get bool value of API limit reached (for apiName API).
// Returns true, if API limit is reached.
func (c cacheRepo) IsAPILimitReached(apiName string) (bool, error) {
	expiredValue, err := c.cacheStorage.Get(apiLimitPrefix + apiName)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to get api limit value", err)
	}
	return len(expiredValue) > 0, nil
}
