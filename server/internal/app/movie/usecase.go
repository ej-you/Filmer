package movie

import (
	"Filmer/server/internal/app/entity"
)

type Usecase interface {
	SearchMovies(searchedMovies *entity.SearchedMovies) error

	CheckMovieExists(movie *entity.Movie) (bool, error)
	GetMovieByKinopoiskID(movie *entity.Movie) (bool, error)
	SaveMovie(movie *entity.Movie) error
}
