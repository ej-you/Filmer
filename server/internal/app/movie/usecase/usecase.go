package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	httpError "Filmer/server/internal/pkg/http_error"
	"Filmer/server/internal/pkg/logger"

	"Filmer/server/internal/app/movie"
)

const (
	officialAPIName   = "official"
	unofficialAPIName = "unofficial"
)

// movie.Usecase interface implementation
type movieUsecase struct {
	cfg                      *config.Config
	logger                   logger.Logger
	movieRepo                movie.Repository
	movieCacheRepo           movie.CacheRepository
	movieKinopoiskWebAPIRepo movie.KinopoiskWebAPIRepository
}

// movie.Usecase constructor
// Returns movie.Usecase interface
func NewUsecase(cfg *config.Config, logger logger.Logger, movieRepo movie.Repository, movieCacheRepo movie.CacheRepository,
	movieKinopoiskWebAPIRepo movie.KinopoiskWebAPIRepository) movie.Usecase {

	return &movieUsecase{
		cfg:                      cfg,
		logger:                   logger,
		movieRepo:                movieRepo,
		movieCacheRepo:           movieCacheRepo,
		movieKinopoiskWebAPIRepo: movieKinopoiskWebAPIRepo,
	}
}

// Search movies by keyword from kinopoisk API
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented
// Fill given searchedMovies struct
func (mu movieUsecase) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	// get searched movies from cache
	found, err := mu.movieCacheRepo.GetSearchMovies(searchedMovies)
	if found {
		return nil
	}
	if err != nil {
		mu.logger.Errorf("movieUsecase.SearchMovies: get from cache: %v", err)
	}

	// check API limit is exhausted
	isLimitExhausted, err := mu.movieCacheRepo.IsAPILimitExhausted(officialAPIName)
	if err != nil {
		return fmt.Errorf("movieUsecase.SearchMovies: %w", err)
	}
	// if daily API limit is exhausted
	if isLimitExhausted {
		return httpError.NewHTTPError(
			http.StatusPaymentRequired,
			"daily API limit is exhausted",
			fmt.Errorf("movieUsecase.SearchMovies: %s API limit", officialAPIName),
		)
	}

	// get searched movies from API
	if err = mu.movieKinopoiskWebAPIRepo.SearchMovies(searchedMovies); err != nil {
		// set API limit to cache if gotten error is 402 error
		return fmt.Errorf("movieUsecase.SearchMovies: %w", mu.setAPILimitIfPaymentError(err, officialAPIName))
	}

	// save searched movies to cache
	err = mu.movieCacheRepo.SetSearchMovies(searchedMovies)
	if err != nil {
		mu.logger.Errorf("movieUsecase.SearchMovies: set to cache: %v", err)
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

	// check API limit is exhausted
	isLimitExhausted, err := mu.movieCacheRepo.IsAPILimitExhausted(unofficialAPIName)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}

	if foundMovie && !isLimitExhausted {
		// run background data update if DB data is outdated AND daily API limit is NOT exhausted
		expiredAt := movie.UpdatedAt.Add(mu.cfg.KinopoiskAPI.DataExpired).UTC()
		now := time.Now().UTC()
		if now.After(expiredAt) {
			go mu.updateMovieData(movie)
		}
	}
	if foundMovie {
		return true, nil
	}

	// if movie was not found in DB and daily API limit is exhausted
	if isLimitExhausted {
		return false, httpError.NewHTTPError(
			http.StatusPaymentRequired,
			"daily API limit is exhausted",
			fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %s API limit", unofficialAPIName),
		)
	}
	// get movie from kinopoisk API if movie was not found in DB
	if err = mu.movieKinopoiskWebAPIRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", mu.setAPILimitIfPaymentError(err, unofficialAPIName))
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
	if err = mu.movieRepo.FullUpdateMovie(movie); err != nil {
		mu.logger.Errorf("background update for film %d: failed to save movie: %v", movie.KinopoiskID, err)
		return
	}
	mu.logger.Infof("Background update for film %d", movie.KinopoiskID)
}

// Set api limit value to cache if given error is 402 payment required error
// Returns cache error if it occurs, else given error
func (mu movieUsecase) setAPILimitIfPaymentError(err error, apiName string) error {
	var httpErr httpError.HTTPError
	// assert gotten error to http error AND check gotten http error is 402
	if errors.As(err, &httpErr) && httpErr.StatusCode() == http.StatusPaymentRequired {
		// set limit is exhausted to cache
		if err := mu.movieCacheRepo.SetAPILimit(apiName); err != nil {
			return err
		}
	}
	return err
}
