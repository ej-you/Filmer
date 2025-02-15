package kinopoisk_api

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)


const getFilmStaffUrl = apiUrl + "/v1/staff"


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
	Producers 	[]Person `json:"producers"`
	Actors 		[]Person `json:"actors"`
}


// получение информации о персонале фильма по его ID
func GetFilmStaff(filmID int, outStruct *FilmStaff) error {
	// создание запроса
	req, err := http.NewRequest("GET", getFilmStaffUrl, nil)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to send request to %q: %v", getFilmStaffUrl, err))
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
	queryParams.Add("filmId", fmt.Sprint(filmID))
	req.URL.RawQuery = queryParams.Encode()

	var rawFilmStaffSlice RawFilmStaffSlice
	// отправка запроса и обработка ответа
	if err := sendRequest(req, getFilmStaffUrl, &rawFilmStaffSlice); err != nil {
		return err
	}

	// инициализация срезов для персонала
	outStruct.Directors = make([]Person, 0)
	outStruct.Producers = make([]Person, 0)
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
			case "PRODUCER":
				outStruct.Producers = append(outStruct.Producers, Person{
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
