package repository

import (
	"fmt"
	"net/http"
	"time"

	"Filmer/server/internal/entity"
	"Filmer/server/internal/movie"
	"Filmer/server/pkg/cache"
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/utils"
)

const (
	apiLimitPrefix     = "api-limit-exhausted:" // key prefix for APIs limit values
	searchMoviesPrefix = "search-movies-data:"  // key prefix for search movies JSON-data
)

// movie.CacheRepository interface implementation
type movieCacheRepository struct {
	cacheClient cache.Cache
	jsonify     jsonify.JSONify
}

// movie.CacheRepository constructor
// Returns movie.CacheRepository interface
func NewCacheRepository(cacheClient cache.Cache, jsonify jsonify.JSONify) movie.CacheRepository {
	return &movieCacheRepository{
		cacheClient: cacheClient,
		jsonify:     jsonify,
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

// Save search movie JSON-data to cache expiring at the beginning of the next day (00:00 UTC)
func (mcr movieCacheRepository) SetSearchMovies(searchedMovies *entity.SearchedMovies) error {
	key := fmt.Sprintf("%s%s:%d", searchMoviesPrefix, searchedMovies.Query, searchedMovies.Page)

	// serialize struct to JSON-data bytes
	bytes, err := mcr.jsonify.Marshal(searchedMovies)
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to serialize search movies data", err)
	}
	// set cache key-value
	err = mcr.cacheClient.Set(key, bytes, utils.ToNextDayDuration(time.Now().UTC()))
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to set search movies data", err)
	}
	return nil
}

// Get search movie JSON-data from cache
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented
// Fill given searchedMovies struct
// Returns true, if search movie JSON-data was found in cache
func (mcr movieCacheRepository) GetSearchMovies(searchedMovies *entity.SearchedMovies) (bool, error) {
	key := fmt.Sprintf("%s%s:%d", searchMoviesPrefix, searchedMovies.Query, searchedMovies.Page)

	// get bytes from cache
	bytesData, err := mcr.cacheClient.GetBytes(key)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to get search movies data", err)
	}
	// if data was not found in cache
	if bytesData == nil {
		return false, nil
	}

	// deserialize JSON-data to struct
	err = mcr.jsonify.Unmarshal(bytesData, searchedMovies)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to deserialize search movies data", err)
	}
	return true, nil
}
