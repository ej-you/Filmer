package kinopoisk_api

import (
	"fmt"
	"server/settings"
)


// структуры для парсинга ответа от API
//easyjson:json
type GenreUnofficial struct {
	Genre string `json:"genre"`
}
//easyjson:json
type RawFilmInfo struct {
	KinopoiskID		int `json:"kinopoiskId"`
	Title			string `json:"nameRu"`
	posterURL 		string `json:"posterUrlPreview"`
	RatingKinopoisk	float64 `json:"ratingKinopoisk"`
	WebUrl 			string `json:"webUrl"`
	Year			int `json:"year"`
	FilmLenMinutes	int `json:"filmLength"`
	Description		string `json:"description"`
	Type			string `json:"type"`
	Genres 			[]GenreUnofficial `json:"genres"`
}

// структура для обработки и возврата ответа, полученного от API
//easyjson:json
// type FilmInfo struct {
// 	ID			int `json:"id"`
// 	Title		string `json:"title"`
// 	ImgUrl		string `json:"imgUrl"`
// 	Rating		string `json:"rating"`
// 	WebUrl 		string `json:"webUrl"`
// 	Year		string `json:"year"`
// 	FilmLength	string `json:"filmLength"`
// 	Description	string `json:"description"`
// 	Type		string `json:"type"`
// 	Genres 		[]string `json:"genres"`
// }


// получение информации о фильме по его ID
func GetFilmInfo(filmID int, outStruct *RawFilmInfo) error {
	newAPI := apiGetRequest{
		URL: fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", filmID),
		APIKey: settings.KinopoiskApiUnofficialKey,
	}
	// отправка запроса и обработка ответа
	err := newAPI.sendRequest(outStruct)
	if err != nil {
		return fmt.Errorf("request to %q: %w", newAPI.URL, err)
	}
	return nil

	// // заполнение обработанной структуры
	// outStruct.ID = rawFilmInfo.KinopoiskID
	// outStruct.Title = rawFilmInfo.Title
	// outStruct.ImgUrl = rawFilmInfo.ImgUrl
	// outStruct.Rating = fmt.Sprint(rawFilmInfo.RatingKinopoisk)
	// outStruct.WebUrl = rawFilmInfo.WebUrl
	// outStruct.Year = fmt.Sprint(rawFilmInfo.Year)
	// outStruct.FilmLength = rawMinutesToTime(rawFilmInfo.FilmLenMinutes)
	// outStruct.Description = rawFilmInfo.Description
	// outStruct.Type = rawFilmInfo.Type
	// outStruct.Genres = rawGenresToSlice(rawFilmInfo.Genres)
	// return nil
}
