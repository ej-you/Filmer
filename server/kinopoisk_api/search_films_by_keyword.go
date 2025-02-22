package kinopoisk_api

import (
	"fmt"
	"server/settings"
)


const filmsPerPage = "25" // кол-во фильмов на странице


// структуры для парсинга ответа от API
//easyjson:json
type Genre struct {
	Genre string `json:"name"`
}
//easyjson:json
type Poster struct {
	URL string `json:"url"`
}
//easyjson:json
type Rating struct {
	Kinopoisk float64 `json:"kp"`
}
//easyjson:json
type Film struct {
	ID		int `json:"id"`
	Title	string `json:"name"`
	Type	string `json:"type"`
	Year	int `json:"year"`
	Genres	[]Genre `json:"genres"`
	Poster	Poster `json:"poster"`
	Rating	Rating `json:"rating"`
}
//easyjson:json
type SearchedFilms struct {
	Films	[]Film `json:"docs"`
	Total	int `json:"total"`
	Limit	int `json:"limit"`
	Page	int `json:"page"`
	Pages	int `json:"pages"`
}


// получение списка фильмов по ключевому слову
func SearchFilmsByKeyword(query string, page int, outStruct *SearchedFilms) error {
	newAPI := apiGetRequest{
		URL: "https://api.kinopoisk.dev/v1.4/movie/search",
		APIKey: settings.KinopoiskApiKey,
		QueryParams: map[string]string{
			"query": query,
			"page": fmt.Sprint(page),
			"limit": filmsPerPage,
		},
	}
	// отправка запроса и обработка ответа
	err := newAPI.sendRequest(outStruct)
	if err != nil {
		return fmt.Errorf("request to %q: %w", newAPI.URL, err)
	}
	return nil
}
