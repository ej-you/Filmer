package entity

import (
	"time"

	"github.com/google/uuid"
)

// movie model
// @description received data about movie
//
//easyjson:json
type Movie struct {
	// movie uuid
	ID uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"id" example:"0a54d171-7b07-41d1-bf5e-0ebe2b64b65a"`
	// movie kinopoisk ID
	KinopoiskID int `gorm:"not null;type:BIGINT;uniqueIndex:uni_movies_kinopoisk_id" json:"kinopoiskID" example:"301"`
	// movie title
	Title string `gorm:"not null;type:VARCHAR(100)" json:"title" example:"Матрица"`
	// movie img URL
	ImgURL string `gorm:"not null;type:VARCHAR(255)" json:"imgURL" example:"https://kinopoiskapiunofficial.tech/images/posters/kp_small/301.jpg"`
	// link to the movie page on Kinopoisk
	WebURL string `gorm:"not null;type:VARCHAR(255)" json:"webURL,omitempty" example:"https://www.kinopoisk.ru/film/301/"`
	// movie rating
	Rating float64 `gorm:"null;type:DECIMAL(2,1)" json:"rating" example:"8.5"`
	// movie release year
	Year int `gorm:"null;type:SMALLINT" json:"year" example:"1999"`
	// movie length
	MovieLength string `gorm:"null;type:VARCHAR(10)" json:"movieLength,omitempty" example:"2:16"`
	// movie description
	Description string `gorm:"null;type:STRING" json:"description,omitempty" example:"Жизнь Томаса Андерсона разделена на две части: днём он — самый обычный офисный работник, получающий нагоняи от начальства, а ночью превращается в хакера по имени Нео, и нет места в сети, куда он бы не смог проникнуть. Но однажды всё меняется. Томас узнаёт ужасающую правду о реальности."`
	// movie type (фильм, сериал, видео, мини-сериал)
	Type string `gorm:"null;type:VARCHAR(20)" json:"type" example:"фильм"`
	// movie staff
	Staff *MovieStaff `gorm:"not null;type:JSONB" json:"staff,omitempty"`
	// last movie data update from Kinopoisk API date (RFC3339 format)
	UpdatedAt time.Time `gorm:"not null;type:TIMESTAMP;index:idx_movies_updated_at,sort:desc" json:"updatedAt" example:"2025-02-28T22:00:05.225526936Z"`

	// movie genres
	Genres     []Genre     `gorm:"foreignKey:MovieID" json:"genres"`
	UserMovies []UserMovie `gorm:"foreignKey:MovieID" json:"-"`
}

func (Movie) TableName() string {
	return "movies"
}
