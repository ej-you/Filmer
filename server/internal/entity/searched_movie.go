package entity

// structs for parsing API response
// @description movie genre (for search movies)
//
//easyjson:json
type SearchedMovieGenre struct {
	Genre string `json:"name" example:"боевик"`
}

// @description movie poster URL (for search movies)
//
//easyjson:json
type SearchedMoviePoster struct {
	URL string `json:"url" example:"https://image.openmoviedb.com/kinopoisk-images/4774061/cf1970bc-3f08-4e0e-a095-2fb57c3aa7c6/orig"`
}

// @description movie rating (for search movies)
//
//easyjson:json
type SearchedMovieRating struct {
	Kinopoisk float64 `json:"kp" example:"8.498"`
}

// @description received movie data (for search movies)
//
//easyjson:json
type SearchedMovie struct {
	// movie kinopoisk ID
	ID int `json:"id" example:"301"`
	// movie title
	Title string `json:"name" example:"Матрица"`
	// movie type
	Type string `json:"type" example:"movie"`
	// movie release year
	Year int `json:"year" example:"1999"`
	// movie genres
	Genres []SearchedMovieGenre `json:"genres"`
	// movie poster
	Poster SearchedMoviePoster `json:"poster"`
	// movie rating
	Rating SearchedMovieRating `json:"rating"`
}

// @description received data (for search movies)
//
//easyjson:json
type SearchedMovies struct {
	// movie info
	Movies []SearchedMovie `json:"docs"`
	// total movies found
	Total int `json:"total" example:"300"`
	// movies amount per page
	Limit int `json:"limit" example:"25"`
	// page number
	Page int `json:"page" example:"1"`
	// all pages amount
	Pages int `json:"pages" example:"12"`
	// keyword
	Query string `json:"query" example:"матрица"`
}
