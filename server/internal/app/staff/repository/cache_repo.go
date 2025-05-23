package repository

import (
	"net/http"
	"strconv"
	"time"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/staff"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/utils"
)

const (
	personPrefix = "person-data:" // key prefix for person JSON-data
)

var _ staff.CacheRepo = (*cacheRepo)(nil)

// staff.CacheRepo implementation.
type cacheRepo struct {
	cacheClient cache.Cache
	jsonify     jsonify.JSONify
}

// Returns staff.CacheRepo interface.
func NewCacheRepository(cacheClient cache.Cache, jsonify jsonify.JSONify) staff.CacheRepo {
	return &cacheRepo{
		cacheClient: cacheClient,
		jsonify:     jsonify,
	}
}

// Save person JSON-data to cache expiring at the beginning of the next day (00:00 UTC).
func (r cacheRepo) SetPersonInfo(person *entity.PersonFull) error {
	key := personPrefix + strconv.Itoa(person.ID)

	// serialize struct to JSON-data bytes
	bytes, err := r.jsonify.Marshal(person)
	if err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError, "serialize person data", err)
	}
	// set cache key-value
	err = r.cacheClient.Set(key, bytes, utils.ToNextDayDuration(time.Now().UTC()))
	if err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError, "set person data", err)
	}
	return nil
}

// Get person JSON-data from cache.
// Person ID must be presented.
// Fill given struct.
// Returns true, if person JSON-data was found in cache.
func (r cacheRepo) GetPersonInfo(person *entity.PersonFull) (bool, error) {
	key := personPrefix + strconv.Itoa(person.ID)

	// get bytes from cache
	bytesData, err := r.cacheClient.GetBytes(key)
	if err != nil {
		return false, httperror.NewHTTPError(http.StatusInternalServerError, "get person data", err)
	}
	// if data was not found in cache
	if bytesData == nil {
		return false, nil
	}

	// deserialize JSON-data to struct
	err = r.jsonify.Unmarshal(bytesData, person)
	if err != nil {
		return false, httperror.NewHTTPError(http.StatusInternalServerError, "deserialize person data", err)
	}
	return true, nil
}
