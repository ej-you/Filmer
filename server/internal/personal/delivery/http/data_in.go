package http

// data for getting person info
//
//easyjson:json
type getPersonInfoIn struct {
	// ID личности
	PersonID int `params:"personID" validate:"required"`
}
