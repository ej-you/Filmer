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

// for parsing API response
//
//easyjson:json
type RawMovieStaff struct {
	StaffID       int    `json:"staffId"`
	Name          string `json:"nameRu"`
	Description   string `json:"description"`
	ProfessionKey string `json:"professionKey"`
	ImgURL        string `json:"posterUrl"`
}

//easyjson:json
type RawMovieStaffSlice []RawMovieStaff
