package usecase

import (
	"fmt"
	"net/http"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/app/usermovie"
	"Filmer/server/internal/pkg/httperror"
)

var _ usermovie.Usecase = (*usecase)(nil)

// usermovie.Usecase implementation.
type usecase struct {
	usermovieDBRepo usermovie.DBRepo
	movieUC         movie.Usecase
}

// Returns usermovie.Usecase interface.
func NewUsecase(usermovieDBRepo usermovie.DBRepo, movieUC movie.Usecase) usermovie.Usecase {
	return &usecase{
		usermovieDBRepo: usermovieDBRepo,
		movieUC:         movieUC,
	}
}

// Get user movie (with full movie info) by its kinopoisk ID.
// Must be presented kinopoisk movie ID (userMovie.Movie.KinopoiskID) and
// user ID (userMovie.UserID).
// Fill given userMovie struct.
// Returns true, if movie was found in DB, else false.
func (u usecase) GetUserMovieByKinopoiskID(userMovie *entity.UserMovie) error {
	// use movieUC to get movie info
	foundInDB, err := u.movieUC.GetMovieByKinopoiskID(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.GetUserMovieByKinopoiskID: %w", err)
	}
	// find user movie if movie was found in DB
	if foundInDB {
		userMovie.MovieID = userMovie.Movie.ID
		_, err = u.usermovieDBRepo.GetUserMovie(userMovie)
		if err != nil {
			return fmt.Errorf("userMovieUsecase.GetUserMovieByKinopoiskID: %w", err)
		}
	}
	return nil
}

// Set stared field of user movie to newStared value.
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID).
// Fill given userMovie struct.
func (u usecase) UpdateUserMovieStared(userMovie *entity.UserMovie, newStared bool) error {
	// init movie struct for user movie
	userMovie.Movie = &entity.Movie{ID: userMovie.MovieID}

	// check that movie with given id is exist in DB
	exists, err := u.movieUC.CheckMovieExists(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	// if movie was not found
	if !exists {
		return httperror.NewHTTPError(http.StatusNotFound,
			"movie was not found", fmt.Errorf("change movie stared to %v", newStared))
	}

	// get or create (if not exists) user movie
	err = u.usermovieDBRepo.FindOrCreateUserMovie(userMovie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	// if user movie stared equals to newStared
	if userMovie.Stared == newStared {
		return nil
	}

	// change stared value
	err = u.usermovieDBRepo.UpdateUserMovieStared(userMovie, newStared)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStared: %w", err)
	}
	return nil
}

// Set status field of user movie to newStatus value.
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID).
// Fill given userMovie struct.
func (u usecase) UpdateUserMovieStatus(userMovie *entity.UserMovie, newStatus int8) error {
	// init movie struct for user movie
	userMovie.Movie = &entity.Movie{ID: userMovie.MovieID}

	// check that movie with given id is exist in DB
	exists, err := u.movieUC.CheckMovieExists(userMovie.Movie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	// if movie was not found
	if !exists {
		return httperror.NewHTTPError(http.StatusNotFound,
			"movie was not found", fmt.Errorf("change movie status to %v", newStatus))
	}

	// get or create (if not exists) user movie
	err = u.usermovieDBRepo.FindOrCreateUserMovie(userMovie)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	// if user movie status equals to newStatus
	if userMovie.Status == newStatus {
		return nil
	}

	// change status value
	err = u.usermovieDBRepo.UpdateUserMovieStatus(userMovie, newStatus)
	if err != nil {
		return fmt.Errorf("userMovieUsecase.UpdateUserMovieStatus: %w", err)
	}
	return nil
}

// Get user movies in given category (stared || want || watched).
// Must be presented category (userMoviesCat.Category) and user ID (userMoviesCat.UserID).
// Fill given userMoviesCat struct.
func (u usecase) GetUserMoviesWithCategory(userMoviesCat *entity.UserMoviesWithCategory) error {
	category := userMoviesCat.Category
	if category != "want" && category != "watched" && category != "stared" {
		return httperror.NewHTTPError(http.StatusInternalServerError,
			"invalid movies category", fmt.Errorf("invalid movies category"))
	}
	if err := u.usermovieDBRepo.GetUserMoviesWithCategory(userMoviesCat); err != nil {
		return fmt.Errorf("userMovieUsecase.GetUserMoviesWithCategory: %w", err)
	}
	return nil
}
