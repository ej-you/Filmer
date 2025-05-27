package entity

import (
	"database/sql/driver"
	"fmt"

	easyjson "github.com/mailru/easyjson"
)

// @description one person info
//
//easyjson:json
type Person struct {
	// person kinopoisk ID
	ID int `json:"id" example:"7836"`
	// person name
	Name string `json:"name" example:"Киану Ривз"`
	// person role (if person is actor)
	Role *string `json:"role,ommitempty" example:"Neo"`
	// person img URL
	ImgURL string `json:"imgUrl" example:"https://st.kp.yandex.net/images/actor_iphone/iphone360_7836.jpg"`
}

// @description movie staff info
//
//easyjson:json
type MovieStaff struct {
	// movie directors
	Directors []Person `json:"directors"`
	// movie actors (up to 30)
	Actors []Person `json:"actors"`
}

// Read JSONB from DB and transformit to MovieStaff struct
func (ms *MovieStaff) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan MovieStaff: expected []byte, got %T", value)
	}
	return easyjson.Unmarshal(bytes, ms)
}

// Save MovieStaff as JSONB into DB
func (ms MovieStaff) Value() (driver.Value, error) {
	return easyjson.Marshal(ms)
}

// @description person movie for person full info
//
//easyjson:json
type PersonFullMovie struct {
	// movie kinopoisk ID
	ID int `json:"id" example:"342"`
	// movie title
	Title string `json:"title" example:"Криминальное чтиво"`
	// person role (if person is actor)
	Role string `json:"role,omitempty" example:"Jimmie"`
}

// @description person full info
//
//easyjson:json
type PersonFull struct {
	// person kinopoisk ID
	ID int `json:"id" example:"7640"`
	// person name
	Name string `json:"name" example:"Квентин Тарантино"`
	// person img URL
	ImgURL string `json:"imgURL" example:"https://kinopoiskapiunofficial.tech/images/actor_posters/kp/7640.jpg"`
	// person sex
	Sex string `json:"sex" example:"мужской"`
	// person profession
	Profession string `json:"profession" example:"Актер, Сценарист, Режиссер"`
	// person age
	Age int `json:"age" example:"62"`
	// person birthday
	Birthday string `json:"birthday" example:"1963-03-27"`
	// person death date (can be not set)
	Death string `json:"death,omitempty" example:"1963-03-27"`
	// facts about person
	Facts []string `json:"facts" example:"Полное имя - Квентин Джером Тарантино.,Имеет двух сестёр и одного брата."`
	// person movies were directed by him
	MoviesDirector []PersonFullMovie `json:"moviesDirector"`
	// person movies in those he was an actor
	MoviesActor []PersonFullMovie `json:"moviesActor"`
}
