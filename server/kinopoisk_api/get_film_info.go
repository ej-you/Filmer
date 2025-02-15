package kinopoisk_api

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)


const getFilmInfoUrl = apiUrl + "/v2.2/films/"


// структура для парсинга ответа от API
//easyjson:json
type RawFilmInfo struct {
	KinopoiskID		int `json:"kinopoiskId"`
	Title			string `json:"nameRu"`
	ImgUrl 			string `json:"posterUrlPreview"`
	RatingKinopoisk	float64 `json:"ratingKinopoisk"`
	WebUrl 			string `json:"webUrl"`
	Year			int `json:"year"`
	FilmLenMinutes	int `json:"filmLength"`
	Description		string `json:"description"`
	Type			string `json:"type"`
	Genres 			[]Genre `json:"genres"` // структура в файле search_films_by_keyword.go
}

// структура для обработки и возврата ответа, полученного от API
//easyjson:json
type FilmInfo struct {
	ID			int `json:"id"`
	Title		string `json:"title"`
	ImgUrl		string `json:"imgUrl"`
	Rating		string `json:"rating"`
	WebUrl 		string `json:"webUrl"`
	Year		string `json:"year"`
	FilmLength	string `json:"filmLength"`
	Description	string `json:"description"`
	Type		string `json:"type"`
	Genres 		[]string `json:"genres"`
}


func rawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/60, minutes%60)
}

func rawGenresToSlice(genres []Genre) []string {
	genresSlice := make([]string, len(genres), len(genres))

	for i, genre := range genres {
		genresSlice[i] = genre.Genre
	}
	return genresSlice
}

// получение информации о фильме по его ID
func GetFilmInfo(filmID int, outStruct *FilmInfo) error {
	// создание запроса
	req, err := http.NewRequest("GET", getFilmInfoUrl+fmt.Sprint(filmID), nil)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to send request to %q: %v", getFilmInfoUrl, err))
	}

	var rawFilmInfo RawFilmInfo
	// отправка запроса и обработка ответа
	if err := sendRequest(req, getFilmInfoUrl, &rawFilmInfo); err != nil {
		return err
	}

	// заполнение обработанной структуры
	outStruct.ID = rawFilmInfo.KinopoiskID
	outStruct.Title = rawFilmInfo.Title
	outStruct.ImgUrl = rawFilmInfo.ImgUrl
	outStruct.Rating = fmt.Sprint(rawFilmInfo.RatingKinopoisk)
	outStruct.WebUrl = rawFilmInfo.WebUrl
	outStruct.Year = fmt.Sprint(rawFilmInfo.Year)
	outStruct.FilmLength = rawMinutesToTime(rawFilmInfo.FilmLenMinutes)
	outStruct.Description = rawFilmInfo.Description
	outStruct.Type = rawFilmInfo.Type
	outStruct.Genres = rawGenresToSlice(rawFilmInfo.Genres)
	return nil
}
