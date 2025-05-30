package repository

import (
	"net/http"
	"time"

	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/utils"
)

const (
	apiLimitPrefix     = "api-limit-exhausted:" // key prefix for APIs limit values
	searchMoviesPrefix = "search-movies-data:"  // key prefix for search movies JSON-data
)

var _apiLimitValue = []byte("t") // value for API limit

var _ movie.CacheRepo = (*cacheRepo)(nil)

// movie.CacheRepo implementation.
type cacheRepo struct {
	cacheStorage cache.Storage
	jsonify      jsonify.JSONify
}

// Returns movie.CacheRepo interface.
func NewCacheRepo(cacheStorage cache.Storage, jsonify jsonify.JSONify) movie.CacheRepo {
	return &cacheRepo{
		cacheStorage: cacheStorage,
		jsonify:      jsonify,
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

// Get bool value of API limit exhausted (for apiName API).
// Returns true, if API limit is exhausted.
func (c cacheRepo) IsAPILimitExhausted(apiName string) (bool, error) {
	expiredValue, err := c.cacheStorage.Get(apiLimitPrefix + apiName)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to get api limit value", err)
	}
	return len(expiredValue) > 0, nil
}
