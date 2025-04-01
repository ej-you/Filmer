package repository

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"

	"Filmer/server/internal/movie"
)

// movie.Repository interface implementation
type movieRepository struct {
	dbClient *gorm.DB
}

// movie.Repository constructor
// Returns movie.Repository interface
func NewRepository(dbClient *gorm.DB) movie.Repository {
	return &movieRepository{
		dbClient: dbClient,
	}
}

// Find movie (without getting info about it)
// Must be presented movie ID (movie.ID)
// Returns true, if movie with given ID was found in DB
func (mr movieRepository) CheckMovieExists(movie *entity.Movie) (bool, error) {
	var foundMovie int64

	selectCountResult := mr.dbClient.Table(movie.TableName()).Where("id = ?", movie.ID).Count(&foundMovie)
	if err := selectCountResult.Error; err != nil {
		return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to find movie with given id", err)
	}
	// if movie was not found
	if foundMovie == 0 {
		return false, nil
	}
	return true, nil
}

// Get movie by its kinopoisk ID
// Must be presented kinopoisk movie ID (movie.KinopoiskID)
// Fill given movie struct
// Returns true, if movie was found in DB, else false
func (mr movieRepository) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	selectResult := mr.dbClient.
		Where("kinopoisk_id = ?", movie.KinopoiskID).
		Preload("Genres").
		First(movie)
	if err := selectResult.Error; err != nil {
		// if NOT "Not found" error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, httpError.NewHTTPError(http.StatusInternalServerError, "failed to get movie", err)
		}
		// if movie was not found
		return false, nil
	}
	return true, nil
}

// Create new movie in DB
func (mr movieRepository) SaveMovie(movie *entity.Movie) error {
	// save movie in DB
	createResult := mr.dbClient.Create(movie)
	if err := createResult.Error; err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to save movie", err)
	}
	return nil
}

// Full update existing movie in DB
func (mr movieRepository) FullUpdateMovie(movie *entity.Movie) error {
	// save movie in DB
	updateResult := mr.dbClient.Save(movie)
	if err := updateResult.Error; err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to full update movie", err)
	}
	return nil
}
