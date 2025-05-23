package movie

import (
	"Filmer/server/internal/entity"
)

type Repository interface {
	CheckMovieExists(movie *entity.Movie) (bool, error)

	GetMovieByKinopoiskID(movie *entity.Movie) (bool, error)
	SaveMovie(movie *entity.Movie) error
	FullUpdateMovie(movie *entity.Movie) error
}

type CacheRepository interface {
	SetAPILimit(apiName string) error
	IsAPILimitExhausted(apiName string) (bool, error)
	SetSearchMovies(searchedMovies *entity.SearchedMovies) error
	GetSearchMovies(searchedMovies *entity.SearchedMovies) (bool, error)
}

type KinopoiskWebAPIRepository interface {
	SearchMovies(searchedMovies *entity.SearchedMovies) error
	GetFullMovieByKinopoiskID(movie *entity.Movie) error
}
