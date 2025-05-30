package movie

import (
	"Filmer/server/internal/app/entity"
)

type DBRepo interface {
	CheckMovieExists(movie *entity.Movie) (bool, error)

	GetMovieByKinopoiskID(movie *entity.Movie) (bool, error)
	SaveMovie(movie *entity.Movie) error
	FullUpdateMovie(movie *entity.Movie) error
}

type CacheRepo interface {
	SetAPILimit(apiName string) error
	IsAPILimitExhausted(apiName string) (bool, error)
	// SetSearchMovies(searchedMovies *entity.SearchedMovies) error
	// GetSearchMovies(searchedMovies *entity.SearchedMovies) (bool, error)
}

type KinopoiskRepo interface {
	SearchMovies(searchedMovies *entity.SearchedMovies) error
	GetFullMovieByKinopoiskID(movie *entity.Movie) error
}
