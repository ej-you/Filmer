package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/kinopoisk"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/app/movie/adapter/amqp"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/utils"
)

var _ movie.Usecase = (*usecase)(nil)

// movie.Usecase implementation.
type usecase struct {
	cfg                *config.Config
	logger             logger.Logger
	movieDBRepo        movie.DBRepo
	movieKinopoiskRepo movie.KinopoiskRepo
	movieAMQPAdapter   *amqp.MovieAdapter
	kinopoiskUC        kinopoisk.Usecase
}

// Returns movie.Usecase interface.
func NewUsecase(cfg *config.Config, logger logger.Logger, movieDBRepo movie.DBRepo,
	movieKinopoiskRepo movie.KinopoiskRepo, movieAMQPAdapter *amqp.MovieAdapter,
	kinopoiskUC kinopoisk.Usecase) movie.Usecase {

	return &usecase{
		cfg:                cfg,
		logger:             logger,
		movieDBRepo:        movieDBRepo,
		movieKinopoiskRepo: movieKinopoiskRepo,
		movieAMQPAdapter:   movieAMQPAdapter,
		kinopoiskUC:        kinopoiskUC,
	}
}

// Search movies by keyword from kinopoisk API
// Query (searchedMovies.Query) and page (searchedMovies.Page) must be presented
// Fill given searchedMovies struct
func (u *usecase) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	// check API limit is reached
	isLimitReached, err := u.kinopoiskUC.IsOfficialAPILimitReached()
	if err != nil {
		return fmt.Errorf("movieUsecase.SearchMovies: %w", err)
	}
	// if daily API limit is reached
	if isLimitReached {
		return httperror.New(http.StatusPaymentRequired, "daily API limit is reached",
			errors.New("movieUsecase.SearchMovies: api limit"),
		)
	}

	// get searched movies from API
	if err := u.movieKinopoiskRepo.SearchMovies(searchedMovies); err != nil {
		// set API limit to cache if gotten error is 402 error
		return fmt.Errorf("movieUsecase.SearchMovies: %w",
			utils.HandleOfficialAPIError(err, u.kinopoiskUC))
	}
	return nil
}

// Find movie (without getting info about it)
// Movie ID must be presented (movie.ID)
// Returns true, if movie with given ID was found in DB
func (u *usecase) CheckMovieExists(movie *entity.Movie) (bool, error) {
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
func (u *usecase) GetMovieByKinopoiskID(movie *entity.Movie) (bool, error) {
	// get movie from DB
	foundMovie, err := u.movieDBRepo.GetMovieByKinopoiskID(movie)
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}

	// check API limit is reached
	isLimitReached, err := u.kinopoiskUC.IsUnofficialAPILimitReached()
	if err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}

	if foundMovie && !isLimitReached {
		u.backgroundUpdate(movie)
	}
	if foundMovie {
		return true, nil
	}

	// if movie was not found in DB and daily API limit is reached
	if isLimitReached {
		return false, httperror.New(http.StatusPaymentRequired, "daily API limit is reached",
			errors.New("movieUsecase.GetMovieByKinopoiskID: api limit"),
		)
	}
	// get movie from kinopoisk API if movie was not found in DB
	if err = u.movieKinopoiskRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w",
			utils.HandleUnofficialAPIError(err, u.kinopoiskUC))
	}
	// save received movie from API in the DB
	if err = u.movieDBRepo.SaveMovie(movie); err != nil {
		return false, fmt.Errorf("movieUsecase.GetMovieByKinopoiskID: %w", err)
	}
	return false, nil
}

// FullUpdate updates movie info.
// Movie ID must be presented.
func (u *usecase) FullUpdate(movie *entity.Movie) error {
	// get kinopoisk ID
	if err := u.movieDBRepo.GetKinopoiskID(movie); err != nil {
		return fmt.Errorf("get movie kinopoisk id: %w", err)
	}
	// request to the kinopoisk API
	if err := u.movieKinopoiskRepo.GetFullMovieByKinopoiskID(movie); err != nil {
		return fmt.Errorf("get movie info from api: %w", err)
	}
	// update movie in DB
	if err := u.movieDBRepo.FullUpdateMovie(movie); err != nil {
		return fmt.Errorf("save movie: %w", err)
	}
	return nil
}

// backgroundUpdate checks that movie data is outdated.
// If so, it sends message with movie ID to RabbitMQ
// for background update by another service.
func (u *usecase) backgroundUpdate(movie *entity.Movie) {
	// if DB data is outdated AND daily API limit is NOT reached
	expiredAt := movie.UpdatedAt.Add(u.cfg.KinopoiskAPI.DataExpired).UTC()
	now := time.Now().UTC()
	if now.After(expiredAt) {
		// send message with movie ID to RabbitMQ
		u.logger.Infof("Run %s movie background update", movie.ID.String())
		if err := u.movieAMQPAdapter.SendID(movie.ID); err != nil {
			u.logger.Errorf("Send message for bg update of %s movie: %v", movie.ID.String(), err)
		}
	}
}
