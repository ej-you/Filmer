package repository

import (
	"net/http"
	"strconv"
	"time"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/personal"
	"Filmer/server/internal/pkg/cache"
	httpError "Filmer/server/internal/pkg/http_error"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/utils"
)

const (
	personPrefix = "person-data:" // key prefix for person JSON-data
)

var _ personal.CacheRepository = (*personalCacheRepository)(nil)

// personal.CacheRepository implementation.
type personalCacheRepository struct {
	cacheClient cache.Cache
	jsonify     jsonify.JSONify
}

func NewCacheRepository(cacheClient cache.Cache, jsonify jsonify.JSONify) personal.CacheRepository {
	return &personalCacheRepository{
		cacheClient: cacheClient,
		jsonify:     jsonify,
	}
}

// Save person JSON-data to cache expiring at the beginning of the next day (00:00 UTC).
func (p personalCacheRepository) SetPersonInfo(person *entity.PersonFull) error {
	key := personPrefix + strconv.Itoa(person.ID)

	// serialize struct to JSON-data bytes
	bytes, err := p.jsonify.Marshal(person)
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "serialize person data", err)
	}
	// set cache key-value
	err = p.cacheClient.Set(key, bytes, utils.ToNextDayDuration(time.Now().UTC()))
	if err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "set person data", err)
	}
	return nil
}

// Get person JSON-data from cache.
// Person ID must be presented.
// Fill given struct.
// Returns true, if person JSON-data was found in cache.
func (p personalCacheRepository) GetPersonInfo(person *entity.PersonFull) (bool, error) {
	key := personPrefix + strconv.Itoa(person.ID)

	// get bytes from cache
	bytesData, err := p.cacheClient.GetBytes(key)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "get person data", err)
	}
	// if data was not found in cache
	if bytesData == nil {
		return false, nil
	}

	// deserialize JSON-data to struct
	err = p.jsonify.Unmarshal(bytesData, person)
	if err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "deserialize person data", err)
	}
	return true, nil
}
