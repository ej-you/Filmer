package movie

import (
	"Filmer/server/internal/app/entity"
)

type DBRepo interface {
	CheckMovieExists(movie *entity.Movie) (bool, error)
	GetKinopoiskID(movie *entity.Movie) error

	GetMovieByKinopoiskID(movie *entity.Movie) (bool, error)
	SaveMovie(movie *entity.Movie) error
	FullUpdateMovie(movie *entity.Movie) error
}

type KinopoiskRepo interface {
	SearchMovies(searchedMovies *entity.SearchedMovies) error
	GetFullMovieByKinopoiskID(movie *entity.Movie) error
}
