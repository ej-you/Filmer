package kinopoisk_api

import (
	"fmt"
	"server/settings"
)


const filmsPerPage = "25" // кол-во фильмов на странице


// структуры для парсинга ответа от API
//easyjson:json
// @description жанр фильма (при поиске фильмов)
type Genre struct {
	Genre string `json:"name" example:"боевик"`
}
//easyjson:json
// @description ссылка на постер фильма (при поиске фильмов)
type Poster struct {
	URL string `json:"url" example:"https://image.openmoviedb.com/kinopoisk-images/4774061/cf1970bc-3f08-4e0e-a095-2fb57c3aa7c6/orig"`
}
//easyjson:json
// @description рейтинг фильма (при поиске фильмов)
type Rating struct {
	Kinopoisk float64 `json:"kp" example:"8.498"`
}
//easyjson:json
// @description получаемые данные о фильме (при поиске фильмов)
type Film struct {
	// kinopoisk ID фильма
	ID		int `json:"id" example:"301"`
	// название фильма
	Title	string `json:"name" example:"Матрица"`
	// тип фильма
	Type	string `json:"type" example:"movie"`
	// год выхода фильма
	Year	int `json:"year" example:"1999"`
	// жанры фильма
	Genres	[]Genre `json:"genres"`
	// постер фильма
	Poster	Poster `json:"poster"`
	// рейтинг фильма
	Rating	Rating `json:"rating"`
}
//easyjson:json
// @description получаемые данные (при поиске фильмов)
type SearchedFilms struct {
	// информация о фильме
	Films	[]Film `json:"docs"`
	// всего найдено результатов
	Total	int `json:"total" example:"300"`
	// количество фильмов на каждой странице
	Limit	int `json:"limit" example:"25"`
	// номер страницы
	Page	int `json:"page" example:"1"`
	// всего страниц
	Pages	int `json:"pages" example:"12"`
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
