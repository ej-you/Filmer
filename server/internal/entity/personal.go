package entity

import (
	"database/sql/driver"
	"fmt"

	easyjson "github.com/mailru/easyjson"
)


//easyjson:json
// @description one person info
type Person struct {
	// person kinopoisk ID
	ID		int 	`json:"id" example:"7836"`
	// person name
	Name	string 	`json:"name" example:"Киану Ривз"`
	// person role (if person is actor)
	Role	*string	`json:"role,ommitempty" example:"Neo"`
	// person img URL
	ImgUrl 	string 	`json:"imgUrl" example:"https://st.kp.yandex.net/images/actor_iphone/iphone360_7836.jpg"`
}

//easyjson:json
// @description movie staff info
type MovieStaff struct {
	// movie directors
	Directors 	[]Person `json:"directors"`
	// movie actors (up to 30)
	Actors 		[]Person `json:"actors"`
}

// Read JSONB from DB and transformit to MovieStaff struct
func (this *MovieStaff) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan MovieStaff: expected []byte, got %T", value)
	}
	return easyjson.Unmarshal(bytes, this)
}

// Save MovieStaff as JSONB into DB
func (this MovieStaff) Value() (driver.Value, error) {
	return easyjson.Marshal(this)
}


// for parsing API response
//easyjson:json
type RawMovieStaff struct {
	StaffID			int		`json:"staffId"`
	Name			string	`json:"nameRu"`
	Description		string	`json:"description"`
	ProfessionKey	string	`json:"professionKey"`
	ImgUrl 			string	`json:"posterUrl"`
}

//easyjson:json
type RawMovieStaffSlice []RawMovieStaff
