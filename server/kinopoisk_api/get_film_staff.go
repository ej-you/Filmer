package kinopoisk_api

import (
	"fmt"
	"server/settings"
)


// структуры для парсинга ответа от API
//easyjson:json
type RawFilmStaffSlice []RawFilmStaff
//easyjson:json
type RawFilmStaff struct {
	StaffID			int `json:"staffId"`
	Name			string `json:"nameRu"`
	Description		string `json:"description"`
	ProfessionKey	string `json:"professionKey"`
	ImgUrl 			string `json:"posterUrl"`
}

// структуры для обработки и возврата ответа, полученного от API
//easyjson:json
type Person struct {
	ID		int `json:"id"`
	Name	string `json:"name"`
	Role	*string `json:"role,ommitempty"`
	ImgUrl 	string `json:"imgUrl"`
}
//easyjson:json
type FilmStaff struct {
	Directors 	[]Person `json:"directors"`
	Actors 		[]Person `json:"actors"`
}


// получение информации о персонале фильма по его ID
func GetFilmStaff(filmID int, outStruct *FilmStaff) error {
	newAPI := apiGetRequest{
		URL: "https://kinopoiskapiunofficial.tech/api/v1/staff",
		APIKey: settings.KinopoiskApiUnofficialKey,
		QueryParams: map[string]string{
			"filmId": fmt.Sprint(filmID),
		},
	}

	var rawFilmStaffSlice RawFilmStaffSlice
	// отправка запроса и обработка ответа
	err := newAPI.sendRequest(&rawFilmStaffSlice)
	if err != nil {
		return fmt.Errorf("request to %q: %w", newAPI.URL, err)
	}
	return nil

	// инициализация срезов для персонала
	outStruct.Directors = make([]Person, 0)
	outStruct.Actors = make([]Person, 0, 30)
	// сортировка персонала по срезам
	for _, rawFilmStaff := range rawFilmStaffSlice {
		switch rawFilmStaff.ProfessionKey {
			case "DIRECTOR":
				outStruct.Directors = append(outStruct.Directors, Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
			case "ACTOR":
				// ограничение на 30 актёров в срезе
				if len(outStruct.Actors) == 30 {
					continue
				}
				outStruct.Actors = append(outStruct.Actors, Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					Role: &rawFilmStaff.Description,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
		}
	}
	return nil
}
