package http

import "github.com/google/uuid"

// data for searching movies with keyword
//
//easyjson:json
type searchFilmsIn struct {
	// keyword for searching
	Query string `query:"q" validate:"required" example:"матрица"`
	// page number
	Page int `query:"page" validate:"required,min=1" example:"1"`
}

// data for update movie info
//
//easyjson:json
type updateMovieIn struct {
	// movie ID
	MovieID uuid.UUID `params:"movieID" validate:"required" example:"86ae41a4-612a-4157-ba82-405872d1d264"`
}
