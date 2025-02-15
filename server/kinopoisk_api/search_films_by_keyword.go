package kinopoisk_api

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)


const searchFilmsByKeywordUrl = apiUrl + "/v2.1/films/search-by-keyword"


// структуры для парсинга ответа от API
//easyjson:json
type Genre struct {
	Genre string `json:"genre"`
}
//easyjson:json
type Film struct {
	FilmID		int `json:"filmId"`
	Title		string `json:"nameRu"`
	Type		string `json:"type"`
	Year		string `json:"year"`
	Genres 		[]Genre `json:"genres"`
	Rating		string `json:"rating"`
	ImgUrl 		string `json:"posterUrlPreview"`
}
//easyjson:json
type SearchedFilms struct {
	SearchFilmsCountResult	int `json:"searchFilmsCountResult"`
	PagesCount				int `json:"pagesCount"`
	Films					[]Film `json:"films"`
}


// получение списка фильмов по ключевому слову
func SearchFilmsByKeyword(keyword string, page int, outStruct *SearchedFilms) error {
	// создание запроса
	req, err := http.NewRequest("GET", searchFilmsByKeywordUrl, nil)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to send request to %q: %v", searchFilmsByKeywordUrl, err))
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
	queryParams.Add("keyword", keyword)
	queryParams.Add("page", fmt.Sprint(page))
	req.URL.RawQuery = queryParams.Encode()

	// отправка запроса и обработка ответа
	return sendRequest(req, searchFilmsByKeywordUrl, outStruct)
}
