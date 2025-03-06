package kinopoisk_api

import (
	"fmt"

	"server/settings"
	"server/db/schemas"
)


// словарь для преобразования типов фильмов
var filmTypesMap = map[string]string {
	"FILM": "фильм",
	"TV_SERIES": "сериал",
	"VIDEO": "видео",
	"MINI_SERIES": "мини-сериал",
	"TV_SHOW": "сериал",
}


// перевод минут в часы:минуты
func rawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/60, minutes%60)
}


// получение информации о фильме по его ID
func GetFilmInfo(filmID int, outStruct *schemas.Movie) error {
	newAPI := apiGetRequest{
		URL: fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", filmID),
		APIKey: settings.KinopoiskApiUnofficialKey,
	}
	// отправка запроса и обработка ответа
	var rawFilmInfo schemas.RawFilmInfo
	err := newAPI.sendRequest(&rawFilmInfo)
	if err != nil {
		return fmt.Errorf("request to %q: %w", newAPI.URL, err)
	}

	// обработка полученных данных
	outStruct.KinopoiskID = rawFilmInfo.KinopoiskID
	outStruct.Title = rawFilmInfo.Title
	outStruct.ImgURL = rawFilmInfo.PosterURL
	outStruct.Rating = rawFilmInfo.RatingKinopoisk
	outStruct.WebURL = rawFilmInfo.WebURL
	outStruct.Year = rawFilmInfo.Year
	// если продолжительность фильма не найдена
	if rawFilmInfo.FilmLenMinutes == 0 {
		outStruct.FilmLength = ""
	} else {
		outStruct.FilmLength = rawMinutesToTime(rawFilmInfo.FilmLenMinutes)
	}
	outStruct.Description = rawFilmInfo.Description
	// по умолчанию - "фильм"
	outStruct.Type = filmTypesMap[rawFilmInfo.Type]
	if outStruct.Type == "" {
		outStruct.Type = "фильм"
	}
	outStruct.Genres = rawFilmInfo.Genres
	return nil
}
