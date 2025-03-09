package repository

import (
	"errors"

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
// Returns true, if movie with given ID was found in DB
func (this movieRepository) CheckMovieExists(movie *entity.Movie) (bool, error) {
	var foundMovie int64

	selectCountResult := this.dbClient.Table(movie.TableName()).Where("id = ?", movie.ID).Count(&foundMovie)
	if err := selectCountResult.Error; err != nil {
		return false, httpError.NewHTTPError(500, "failed to find movie with given id", err)
	}
	// if movie was not found
	if foundMovie == 0 {
		return false, nil
	}
	return true, nil
}

// Get movie by its kinopoisk ID
// Fill given movie struct
// Returns true, if movie was found in DB, else false
func (this movieRepository) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	selectResult := this.dbClient.
		Where("kinopoisk_id = ?", movie.KinopoiskID).
		Preload("Genres").
		First(movie)
	if err := selectResult.Error; err != nil {
		// if NOT "Not found" error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, httpError.NewHTTPError(500, "failed to get movie", err)
		}
		// if movie was not found
		return false, nil
	}
	return true, nil
}

// Create new movie in DB
func (this movieRepository) SaveMovie(movie *entity.Movie) error {
	// save movie in DB
	createResult := this.dbClient.Create(movie)
	if err := createResult.Error; err != nil {
		return httpError.NewHTTPError(500, "failed to save movie", err)
	}
	return nil
}
