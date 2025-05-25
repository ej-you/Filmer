package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/logger"
)

const (
	officialAPIName   = "official"
	unofficialAPIName = "unofficial"
)

var _ movie.Usecase = (*usecase)(nil)

// movie.Usecase implementation.
type usecase struct {
	cfg                *config.Config
	logger             logger.Logger
	movieDBRepo        movie.DBRepo
	movieCacheRepo     movie.CacheRepo
	movieKinopoiskRepo movie.KinopoiskRepo
}

// Returns movie.Usecase interface.
func NewUsecase(cfg *config.Config, logger logger.Logger, movieDBRepo movie.DBRepo,
	movieCacheRepo movie.CacheRepo, movieKinopoiskRepo movie.KinopoiskRepo) movie.Usecase {

	return &usecase{
		cfg:                cfg,
		logger:             logger,
		movieDBRepo:        movieDBRepo,
		movieCacheRepo:     movieCacheRepo,
		movieKinopoiskRepo: movieKinopoiskRepo,
	}
}

// Search movies by keyword from kinopoisk API
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented
// Fill given searchedMovies struct
func (u usecase) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	// get searched movies from cache
	found, err := u.movieCacheRepo.GetSearchMovies(searchedMovies)
	if found {
		return nil
	}
	if err != nil {
		u.logger.Errorf("movieUsecase.SearchMovies: get from cache: %v", err)
	}

	// check API limit is exhausted
	isLimitExhausted, err := u.movieCacheRepo.IsAPILimitExhausted(officialAPIName)
	if err != nil {
		return fmt.Errorf("movieUsecase.SearchMovies: %w", err)
	}
	// if daily API limit is exhausted
	if isLimitExhausted {
		return httperror.New(
			http.StatusPaymentRequired,
			"daily API limit is exhausted",
			fmt.Errorf("movieUsecase.SearchMovies: %s API limit", officialAPIName),
		)
	}

	// get searched movies from API
	if err = u.movieKinopoiskRepo.SearchMovies(searchedMovies); err != nil {
		// set API limit to cache if gotten error is 402 error
		return fmt.Errorf("movieUsecase.SearchMovies: %w",
			u.setAPILimitIfPaymentError(err, officialAPIName))
	}

	// save searched movies to cache
	err = u.movieCacheRepo.SetSearchMovies(searchedMovies)
	if err != nil {
		u.logger.Errorf("movieUsecase.SearchMovies: set to cache: %v", err)
	}

	return nil
}

// Find movie (without getting info about it)
// Movie ID must be presented (movie.ID)
// Returns true, if movie with given ID was found in DB
func (u usecase) CheckMovieExists(movie *entity.Movie) (bool, error) {
	exists, err := u.movieDBRepo.CheckMovieExists(movie)
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
func (u usecase) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	// get movie from DB
	foundMovie, err := u.movieDBRepo.GetMovieByKinopoiskID(movie)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}

	// check API limit is exhausted
	isLimitExhausted, err := u.movieCacheRepo.IsAPILimitExhausted(unofficialAPIName)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}

	if foundMovie && !isLimitExhausted {
		// run background data update if DB data is outdated AND daily API limit is NOT exhausted
		expiredAt := movie.UpdatedAt.Add(u.cfg.KinopoiskAPI.DataExpired).UTC()
		now := time.Now().UTC()
		if now.After(expiredAt) {
			go u.updateMovieData(movie)
		}
	}
	if foundMovie {
		return true, nil
	}

	// if movie was not found in DB and daily API limit is exhausted
	if isLimitExhausted {
		return false, httperror.New(
			http.StatusPaymentRequired,
			"daily API limit is exhausted",
			fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %s API limit", unofficialAPIName),
		)
	}
	// get movie from kinopoisk API if movie was not found in DB
	if err = u.movieKinopoiskRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w",
			u.setAPILimitIfPaymentError(err, unofficialAPIName))
	}
	// save received movie from API in the DB
	if err = u.movieDBRepo.SaveMovie(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}
	return false, nil
}

// Save movie to DB
// Must be presented all movie fields
func (u usecase) SaveMovie(movie *entity.Movie) error {
	if err := u.movieDBRepo.SaveMovie(movie); err != nil {
		return fmt.Errorf("movieUsecase.SaveMovie: %w", err)
	}
	return nil
}

// Background movie info update
// Must be presented movie ID (movie.ID) and kinopoisk movie ID (movie.KinopoiskID)
func (u usecase) updateMovieData(movie *entity.Movie) {
	var err error
	// turn to the kinopoisk API
	if err = u.movieKinopoiskRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		u.logger.Errorf("background update for film %d: failed to get film info from API: %v",
			movie.KinopoiskID, err)
		return
	}
	// update movie in DB (so, PK is presented then data will update)
	if err = u.movieDBRepo.FullUpdateMovie(movie); err != nil {
		u.logger.Errorf("background update for film %d: failed to save movie: %v",
			movie.KinopoiskID, err)
		return
	}
	u.logger.Infof("Background update for film %d", movie.KinopoiskID)
}

// Set api limit value to cache if given error is 402 payment required error
// Returns cache error if it occurs, else given error
func (u usecase) setAPILimitIfPaymentError(err error, apiName string) error {
	var httpErr httperror.HTTPError
	// assert gotten error to http error AND check gotten http error is 402
	if errors.As(err, &httpErr) && httpErr.StatusCode() == http.StatusPaymentRequired {
		// set limit is exhausted to cache
		if err := u.movieCacheRepo.SetAPILimit(apiName); err != nil {
			return err
		}
	}
	return err
}
