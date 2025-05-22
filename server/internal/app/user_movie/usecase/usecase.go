package usecase

import (
	"fmt"
	"net/http"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	httpError "Filmer/server/internal/pkg/http_error"

	userMovie "Filmer/server/internal/app/user_movie"
)

// userMovie.Usecase interface implementation
type userMovieUsecase struct {
	userMovieRepo userMovie.Repository
	movieUC       movie.Usecase
}

// userMovie.Usecase constructor
// Returns userMovie.Usecase interface
func NewUsecase(userMovieRepo userMovie.Repository, movieUC movie.Usecase) userMovie.Usecase {
	return &userMovieUsecase{
		userMovieRepo: userMovieRepo,
		movieUC:       movieUC,
	}
}

// Get user movie (with full movie info) by its kinopoisk ID
// Must be presented kinopoisk movie ID (userMovie.Movie.KinopoiskID) and user ID (userMovie.UserID)
// Fill given userMovie struct
// Returns true, if movie was found in DB, else false
func (umu userMovieUsecase) GetUserMovieByKinopoiskID(userMovie *entity.UserMovie) error {
	// use movieUC to get movie info
	foundInDB, err := umu.movieUC.GetMovieByKinopoiskID(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.GetUserMovieByKinopoiskID: %w", err)
	}
	// find user movie if movie was found in DB
	if foundInDB {
		userMovie.MovieID = userMovie.Movie.ID
		_, err = umu.userMovieRepo.GetUserMovie(userMovie)
		if err != nil {
			return fmt.Errorf("userMovieUsecase.GetUserMovieByKinopoiskID: %w", err)
		}
	}
	return nil
}

// Set stared field of user movie to newStared value
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Fill given userMovie struct
func (umu userMovieUsecase) UpdateUserMovieStared(userMovie *entity.UserMovie, newStared bool) error {
	// init movie struct for user movie
	userMovie.Movie = &entity.Movie{ID: userMovie.MovieID}

	// check that movie with given id is exist in DB
	exists, err := umu.movieUC.CheckMovieExists(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	// if movie was not found
	if !exists {
		return httpError.NewHTTPError(http.StatusNotFound, "movie with given id was not found", fmt.Errorf("change movie stared to %v", newStared))
	}

	// get or create (if not exists) user movie
	err = umu.userMovieRepo.FindOrCreateUserMovie(userMovie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	// if user movie stared equals to newStared
	if userMovie.Stared == newStared {
		return nil
	}

	// change stared value
	err = umu.userMovieRepo.UpdateUserMovieStared(userMovie, newStared)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	return nil
}

// Set status field of user movie to newStatus value
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Fill given userMovie struct
func (umu userMovieUsecase) UpdateUserMovieStatus(userMovie *entity.UserMovie, newStatus int8) error {
	// init movie struct for user movie
	userMovie.Movie = &entity.Movie{ID: userMovie.MovieID}

	// check that movie with given id is exist in DB
	exists, err := umu.movieUC.CheckMovieExists(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	// if movie was not found
	if !exists {
		return httpError.NewHTTPError(http.StatusNotFound, "movie with given id was not found", fmt.Errorf("change movie status to %v", newStatus))
	}

	// get or create (if not exists) user movie
	err = umu.userMovieRepo.FindOrCreateUserMovie(userMovie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	// if user movie status equals to newStatus
	if userMovie.Status == newStatus {
		return nil
	}

	// change status value
	err = umu.userMovieRepo.UpdateUserMovieStatus(userMovie, newStatus)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	return nil
}

// Get user movies in given category (stared || want || watched)
// Must be presented category (userMoviesCat.Category) and user ID (userMoviesCat.UserID)
// Fill given userMoviesCat struct
func (umu userMovieUsecase) GetUserMoviesWithCategory(userMoviesCat *entity.UserMoviesWithCategory) error {
	if userMoviesCat.Category != "stared" && userMoviesCat.Category != "want" && userMoviesCat.Category != "watched" {
		return httpError.NewHTTPError(http.StatusInternalServerError, "invalid movies category", fmt.Errorf("invalid movies category"))
	}
	if err := umu.userMovieRepo.GetUserMoviesWithCategory(userMoviesCat); err != nil {
		return fmt.Errorf("userMovieUsecase.GetUserMoviesWithCategory: %w", err)
	}
	return nil
}
