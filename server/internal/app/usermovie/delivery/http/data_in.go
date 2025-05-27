package http

import (
	"github.com/google/uuid"

	"Filmer/server/internal/app/entity"
)

// data for getting movie info
//
//easyjson:json
type getFilmInfoIn struct {
	// kinopoisk ID фильма
	KinopoiskID int `params:"kinopoiskID" validate:"required"`
}

// data for update movie category
//
//easyjson:json
type setFilmCategoryIn struct {
	// movie ID
	MovieID uuid.UUID `params:"movieID" validate:"required" example:"86ae41a4-612a-4157-ba82-405872d1d264"`
}

// data for getting user movie list with filter, sort and pagination
//
//easyjson:json
type categoryFilmsIn struct {
	// filter fields
	entity.UserMoviesFilter
	// sort field
	entity.UserMoviesSort
	// pagination
	entity.UserMoviesPagination
}
