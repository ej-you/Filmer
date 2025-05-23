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
		return false, httperror.NewHTTPError(http.StatusInternalServerError,
			"failed to find movie with given id", err)
	}
	// if movie was not found
	if foundMovie == 0 {
		return false, nil
	}
	return true, nil
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
			return false, httperror.NewHTTPError(http.StatusInternalServerError,
				"failed to get movie", err)
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
		return httperror.NewHTTPError(http.StatusInternalServerError,
			"failed to save movie", err)
	}
	return nil
}

// Full update existing movie in DB.
func (r dbRepo) FullUpdateMovie(movie *entity.Movie) error {
	// save movie in DB
	updateResult := r.dbClient.Save(movie)
	if err := updateResult.Error; err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError,
			"failed to full update movie", err)
	}
	return nil
}
