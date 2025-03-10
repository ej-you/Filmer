package user_movie

import (
	"Filmer/server/internal/entity"
)


type Repository interface {
	GetUserMovie(userMovie *entity.UserMovie) (bool, error)
	FindOrCreateUserMovie(userMovie *entity.UserMovie) error
	
	UpdateUserMovieStared(userMovie *entity.UserMovie, newStared bool) error
	UpdateUserMovieStatus(userMovie *entity.UserMovie, newStatus int8) error
	
	GetUserMoviesWithCategory(userMoviesWithCategory *entity.UserMoviesWithCategory) error
}
