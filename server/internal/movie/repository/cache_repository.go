package repository

import (
	"net/http"
	"time"

	"Filmer/server/pkg/cache"
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/utils"

	"Filmer/server/internal/movie"
)

const apiLimitPrefix = "api-limit-exhausted:" // key prefix for APIs limit values

// movie.CacheRepository interface implementation
type movieCacheRepository struct {
	cacheClient cache.Cache
}

// movie.CacheRepository constructor
// Returns movie.CacheRepository interface
func NewCacheRepository(cacheClient cache.Cache) movie.CacheRepository {
	return &movieCacheRepository{
		cacheClient: cacheClient,
	}
}

// Set api limit (for apiName API) expiring at the beginning of the next day (00:00 UTC)
func (mcr movieCacheRepository) SetAPILimit(apiName string) error {
	err := mcr.cacheClient.Set(apiLimitPrefix+apiName, "true", utils.ToNextDayDuration(time.Now().UTC()))
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to set api limit", err)
	}
	return nil
}

// Get bool value of API limit exhausted (for apiName API)
// Returns true, if API limit is exhausted
func (mcr movieCacheRepository) IsAPILimitExhausted(apiName string) (bool, error) {
	isExhausted, err := mcr.cacheClient.GetBool(apiLimitPrefix + apiName)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to get api limit value", err)
	}
	return isExhausted, nil
}
