package repository

import "Filmer/server/internal/app/entity"

// for parsing API response with movie info
//
//easyjson:json
type rawMovie struct {
	KinopoiskID     int            `json:"kinopoiskId"`
	Title           string         `json:"nameRu"`
	PosterURL       string         `json:"posterUrlPreview"`
	WebURL          string         `json:"webUrl"`
	RatingKinopoisk float64        `json:"ratingKinopoisk"`
	Year            int            `json:"year"`
	FilmLenMinutes  int            `json:"filmLength"`
	Description     string         `json:"description"`
	Type            string         `json:"type"`
	Genres          []entity.Genre `json:"genres"`
}

// for parsing API response with movie staff info
//
//easyjson:json
type rawMoviePerson struct {
	StaffID       int    `json:"staffId"`
	Name          string `json:"nameRu"`
	Description   string `json:"description"`
	ProfessionKey string `json:"professionKey"`
	ImgURL        string `json:"posterUrl"`
}

//easyjson:json
type rawMovieStaff []rawMoviePerson
