package repository

// for parsing API response with full person info
//
//easyjson:json
type rawPersonFull struct {
	PersonID   int                  `json:"personId"`
	Name       string               `json:"nameRu"`
	ImgURL     string               `json:"posterUrl"`
	Sex        string               `json:"sex"`
	Profession string               `json:"profession"`
	Age        int                  `json:"age"`
	Birthday   string               `json:"birthday"`
	Death      string               `json:"death,omitempty"`
	Facts      []string             `json:"facts"`
	Movies     []rawPersonFullMovie `json:"films"`
}

//easyjson:json
type rawPersonFullMovie struct {
	ID            int    `json:"filmId"`
	Name          string `json:"nameRu"`
	Description   string `json:"description,omitempty"`
	Rating        string `json:"rating"`
	ProfessionKey string `json:"professionKey"`
}
