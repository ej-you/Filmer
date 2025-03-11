package movie

import (
	"Filmer/server/internal/entity"
)

type Repository interface {
	CheckMovieExists(movie *entity.Movie) (bool, error)

	GetMovieByKinopoiskID(movie *entity.Movie) (bool, error)
	SaveMovie(movie *entity.Movie) error
}

type KinopoiskWebAPIRepository interface {
	SearchMovies(searchedMovies *entity.SearchedMovies) error
	GetFullMovieByKinopoiskID(movie *entity.Movie) error
}
