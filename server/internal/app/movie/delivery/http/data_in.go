package http

// data for searching movies with keyword
//
//easyjson:json
type searchFilmsIn struct {
	// keyword for searching
	Query string `query:"q" validate:"required" example:"матрица"`
	// page number
	Page int `query:"page" validate:"required,min=1" example:"1"`
}
