package entity

import (
	"database/sql/driver"
	"fmt"

	easyjson "github.com/mailru/easyjson"
)


// информация о персонале фильма
//easyjson:json
// @description информация об одном человеке
type Person struct {
	// kinopoisk ID человека
	ID		int 	`json:"id" example:"7836"`
	// имя человека
	Name	string 	`json:"name" example:"Киану Ривз"`
	// роль (у актёра)
	Role	*string	`json:"role,ommitempty" example:"Neo"`
	// ссылка на картинку человека
	ImgUrl 	string 	`json:"imgUrl" example:"https://st.kp.yandex.net/images/actor_iphone/iphone360_7836.jpg"`
}

//easyjson:json
// @description информация о персонале фильма
type FilmStaff struct {
	// режиссёры фильма
	Directors 	[]Person `json:"directors"`
	// актёры фильма (до 30 максимум)
	Actors 		[]Person `json:"actors"`
}

// метод для чтения JSONB из БД и преобразования в структуру FilmStaff
func (this *FilmStaff) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan FilmStaff: expected []byte, got %T", value)
	}
	return easyjson.Unmarshal(bytes, this)
}

// метод для сохранения FilmStaff в виде JSONB в БД
func (this FilmStaff) Value() (driver.Value, error) {
	return easyjson.Marshal(this)
}


// структуры для парсинга ответа от API
//easyjson:json
type RawFilmStaff struct {
	StaffID			int		`json:"staffId"`
	Name			string	`json:"nameRu"`
	Description		string	`json:"description"`
	ProfessionKey	string	`json:"professionKey"`
	ImgUrl 			string	`json:"posterUrl"`
}

//easyjson:json
type RawFilmStaffSlice []RawFilmStaff
