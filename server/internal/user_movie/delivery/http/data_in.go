package http

import (
	"github.com/google/uuid"

	"Filmer/server/internal/entity"
)


//easyjson:json
// data for getting movie info
type getFilmInfoIn struct {
	// kinopoisk ID фильма
	KinopoiskID int `params:"kinopoiskID" validate:"required"`
}

//easyjson:json
// data for update movie category
type setFilmCategoryIn struct {
	// movie ID
	MovieID uuid.UUID `params:"movieID" validate:"required" example:"86ae41a4-612a-4157-ba82-405872d1d264"`
}

//easyjson:json
// data for getting user movie list with filter, sort and pagination
type categoryFilmsIn struct {
	// filter fields
	entity.UserMoviesFilter
	// sort field
	entity.UserMoviesSort
	// pagination
	entity.UserMoviesPagination
}
