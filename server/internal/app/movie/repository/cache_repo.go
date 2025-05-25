package repository

import (
	"fmt"
	"net/http"
	"time"

	"Filmer/server/internal/app/entity"
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

var _ movie.CacheRepo = (*cacheRepo)(nil)

// movie.CacheRepo implementation.
type cacheRepo struct {
	cacheClient cache.Cache
	jsonify     jsonify.JSONify
}

// Returns movie.CacheRepo interface.
func NewCacheRepo(cacheClient cache.Cache, jsonify jsonify.JSONify) movie.CacheRepo {
	return &cacheRepo{
		cacheClient: cacheClient,
		jsonify:     jsonify,
	}
}

// Set api limit (for apiName API) expiring at the beginning of the next day (00:00 UTC).
func (c cacheRepo) SetAPILimit(apiName string) error {
	toNextDay := utils.ToNextDayDuration(time.Now().UTC())
	err := c.cacheClient.Set(apiLimitPrefix+apiName, "true", toNextDay)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to set api limit", err)
	}
	return nil
}

// Get bool value of API limit exhausted (for apiName API).
// Returns true, if API limit is exhausted.
func (c cacheRepo) IsAPILimitExhausted(apiName string) (bool, error) {
	isExhausted, err := c.cacheClient.GetBool(apiLimitPrefix + apiName)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to get api limit value", err)
	}
	return isExhausted, nil
}

// Save search movie JSON-data to cache expiring at the beginning of the next day (00:00 UTC).
func (c cacheRepo) SetSearchMovies(searchedMovies *entity.SearchedMovies) error {
	key := fmt.Sprintf("%s%s:%d", searchMoviesPrefix, searchedMovies.Query, searchedMovies.Page)

	// serialize struct to JSON-data bytes
	bytes, err := c.jsonify.Marshal(searchedMovies)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to serialize search movies data", err)
	}
	// set cache key-value
	err = c.cacheClient.Set(key, bytes, utils.ToNextDayDuration(time.Now().UTC()))
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to set search movies data", err)
	}
	return nil
}

// Get search movie JSON-data from cache.
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented.
// Fill given searchedMovies struct.
// Returns true, if search movie JSON-data was found in cache.
func (c cacheRepo) GetSearchMovies(searchedMovies *entity.SearchedMovies) (bool, error) {
	key := fmt.Sprintf("%s%s:%d", searchMoviesPrefix, searchedMovies.Query, searchedMovies.Page)

	// get bytes from cache
	bytesData, err := c.cacheClient.GetBytes(key)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to get search movies data", err)
	}
	// if data was not found in cache
	if bytesData == nil {
		return false, nil
	}

	// deserialize JSON-data to struct
	err = c.jsonify.Unmarshal(bytesData, searchedMovies)
	if err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"failed to deserialize search movies data", err)
	}
	return true, nil
}
