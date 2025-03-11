package usecase

import (
	"fmt"
	"time"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	"Filmer/server/pkg/logger"

	"Filmer/server/internal/movie"
)

// movie.Usecase interface implementation
type movieUsecase struct {
	cfg                      *config.Config
	logger                   logger.Logger
	movieRepo                movie.Repository
	movieKinopoiskWebAPIRepo movie.KinopoiskWebAPIRepository
}

// movie.Usecase constructor
// Returns movie.Usecase interface
func NewUsecase(cfg *config.Config, logger logger.Logger, movieRepo movie.Repository,
	movieKinopoiskWebAPIRepo movie.KinopoiskWebAPIRepository) movie.Usecase {

	return &movieUsecase{
		cfg:                      cfg,
		logger:                   logger,
		movieRepo:                movieRepo,
		movieKinopoiskWebAPIRepo: movieKinopoiskWebAPIRepo,
	}
}

// Search movies by keyword from kinopoisk API
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented
// Fill given searchedMovies struct
func (mu movieUsecase) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	if err := mu.movieKinopoiskWebAPIRepo.SearchMovies(searchedMovies); err != nil {
		return fmt.Errorf("movieUsecase.SearchMovies: %w", err)
	}
	return nil
}

// Find movie (without getting info about it)
// Movie ID must be presented (movie.ID)
// Returns true, if movie with given ID was found in DB
func (mu movieUsecase) CheckMovieExists(movie *entity.Movie) (bool, error) {
	exists, err := mu.movieRepo.CheckMovieExists(movie)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.CheckMovieExists: %w", err)
	}
	return exists, nil
}

// Get movie by its kinopoisk ID
// First try to find movie in DB. If there is no in DB then send request to API
// Must be presented kinopoisk movie ID (movie.KinopoiskID)
// Fill given movie struct
// Returns true, if movie was found in DB, else false
func (mu movieUsecase) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	// get movie from DB
	foundMovie, err := mu.movieRepo.GetMovieByKinopoiskID(movie)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}
	if foundMovie {
		// run background data update if DB data is outdated
		expiredAt := movie.UpdatedAt.Add(mu.cfg.KinopoiskAPI.DataExpired).UTC()
		now := time.Now().UTC()
		if now.After(expiredAt) {
			go mu.updateMovieData(movie)
		}

		return true, nil
	}
	// get movie from kinopoisk API if movie was not found in DB
	if err = mu.movieKinopoiskWebAPIRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}
	// save received movie from API in the DB
	if err = mu.movieRepo.SaveMovie(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}
	return false, nil
}

// Save movie to DB
// Must be presented all movie fields
func (mu movieUsecase) SaveMovie(movie *entity.Movie) error {
	if err := mu.movieRepo.SaveMovie(movie); err != nil {
		return fmt.Errorf("movieUsecase.SaveMovie: %w", err)
	}
	return nil
}

// Background movie info update
// Must be presented movie ID (movie.ID) and kinopoisk movie ID (movie.KinopoiskID)
func (mu movieUsecase) updateMovieData(movie *entity.Movie) {
	var err error
	// turn to the kinopoisk API
	if err = mu.movieKinopoiskWebAPIRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		mu.logger.Errorf("background update for film %d: failed to get film info from API: %v", movie.KinopoiskID, err)
		return
	}
	// update movie in DB (so, PK is presented then data will update)
	if err = mu.movieRepo.SaveMovie(movie); err != nil {
		mu.logger.Errorf("background update for film %d: failed to save movie: %v", movie.KinopoiskID, err)
		return
	}
	mu.logger.Infof("Background update for film %d", movie.KinopoiskID)
}
