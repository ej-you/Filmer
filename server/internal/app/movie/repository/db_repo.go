package repository

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/pkg/httperror"
)

var _ movie.DBRepo = (*dbRepo)(nil)

// movie.DBRepo implementation.
type dbRepo struct {
	dbClient *gorm.DB
}

// Returns movie.DBRepo interface.
func NewDBRepo(dbClient *gorm.DB) movie.DBRepo {
	return &dbRepo{
		dbClient: dbClient,
	}
}

// Find movie (without getting info about it).
// Must be presented movie ID (movie.ID).
// Returns true, if movie with given ID was found in DB.
func (r dbRepo) CheckMovieExists(movie *entity.Movie) (bool, error) {
	var foundMovie int64

	selectCountResult := r.dbClient.
		Table(movie.TableName()).
		Where("id = ?", movie.ID).
		Count(&foundMovie)
	if err := selectCountResult.Error; err != nil {
		return false, httperror.New(http.StatusInternalServerError,
			"find movie with given id", err)
	}
	return foundMovie != 0, nil
}

// GetKinopoiskID gets kinopoisk ID for movie by given ID (without getting other info).
// Movie ID must be presented (movie.ID).
// Fill KinopoiskID field of given struct. Returns 404 error if movie was not found.
func (r dbRepo) GetKinopoiskID(movie *entity.Movie) error {
	selectResult := r.dbClient.
		Select("kinopoisk_id").
		Where("id = ?", movie.ID).
		First(movie)
	if err := selectResult.Error; err != nil {
		// if movie was not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httperror.New(http.StatusNotFound,
				"movie not found", err)
		}
		return httperror.New(http.StatusInternalServerError,
			"get movie", err)
	}
	return nil
}

// Get movie by its kinopoisk ID.
// Must be presented kinopoisk movie ID (movie.KinopoiskID).
// Fill given movie struct.
// Returns true, if movie was found in DB, else false.
func (r dbRepo) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	selectResult := r.dbClient.
		Where("kinopoisk_id = ?", movie.KinopoiskID).
		Preload("Genres").
		First(movie)
	if err := selectResult.Error; err != nil {
		// if NOT "Not found" error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, httperror.New(http.StatusInternalServerError,
				"get movie", err)
		}
		// if movie was not found
		return false, nil
	}
	return true, nil
}

// Create new movie in DB.
func (r dbRepo) SaveMovie(movie *entity.Movie) error {
	// save movie in DB
	createResult := r.dbClient.Create(movie)
	if err := createResult.Error; err != nil {
		return httperror.New(http.StatusInternalServerError,
			"save movie", err)
	}
	return nil
}

// Full update existing movie in DB.
// Movie ID must be presented (so, PK is presented then data will update).
func (r dbRepo) FullUpdateMovie(movie *entity.Movie) error {
	// save movie in DB
	updateResult := r.dbClient.Save(movie)
	if err := updateResult.Error; err != nil {
		return httperror.New(http.StatusInternalServerError,
			"full update movie", err)
	}
	return nil
}
