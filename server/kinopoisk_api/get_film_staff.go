package kinopoisk_api

import (
	"fmt"

	"server/settings"
	"server/db/schemas"
)


// получение информации о персонале фильма по его ID
func GetFilmStaff(filmID int, outStruct *schemas.FilmStaff) error {
	newAPI := apiGetRequest{
		URL: "https://kinopoiskapiunofficial.tech/api/v1/staff",
		APIKey: settings.KinopoiskApiUnofficialKey,
		QueryParams: map[string]string{
			"filmId": fmt.Sprint(filmID),
		},
	}

	var rawFilmStaffSlice schemas.RawFilmStaffSlice
	// отправка запроса и обработка ответа
	err := newAPI.sendRequest(&rawFilmStaffSlice)
	if err != nil {
		return fmt.Errorf("request to %q: %w", newAPI.URL, err)
	}

	// инициализация срезов для персонала
	outStruct.Directors = make([]schemas.Person, 0)
	outStruct.Actors = make([]schemas.Person, 0, 30)
	// сортировка персонала по срезам
	for _, rawFilmStaff := range rawFilmStaffSlice {
		switch rawFilmStaff.ProfessionKey {
			case "DIRECTOR":
				outStruct.Directors = append(outStruct.Directors, schemas.Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
			case "ACTOR":
				// ограничение на 30 актёров в срезе
				if len(outStruct.Actors) == 30 {
					continue
				}
				outStruct.Actors = append(outStruct.Actors, schemas.Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					Role: &rawFilmStaff.Description,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
		}
	}
	return nil
}
