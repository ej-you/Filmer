package entity

import (
	"time"

	"github.com/google/uuid"
)


// модель фильма
//easyjson:json
// @description получаемые данные о фильме
type Movie struct {
	// uuid фильма
	ID			uuid.UUID	`gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"id" example:"0a54d171-7b07-41d1-bf5e-0ebe2b64b65a"`
	// kinopoisk ID фильма
	KinopoiskID	int			`gorm:"not null;type:BIGINT;uniqueIndex:uni_movies_kinopoisk_id" json:"kinopoiskID" example:"301"`
	// название фильма
	Title		string		`gorm:"not null;type:VARCHAR(100)" json:"title" example:"Матрица"`
	// ссылка на картинку фильма
	ImgURL 		string		`gorm:"not null;type:VARCHAR(255)" json:"imgURL" example:"https://kinopoiskapiunofficial.tech/images/posters/kp_small/301.jpg"`
	// ссылка на страницу фильма на Kinopoisk
	WebURL 		string		`gorm:"not null;type:VARCHAR(255)" json:"webURL,omitempty" example:"https://www.kinopoisk.ru/film/301/"`
	// рейтинг фильма
	Rating		float64		`gorm:"null;type:DECIMAL(2,1)" json:"rating" example:"8.5"`
	// год выхода фильма
	Year		int			`gorm:"null;type:SMALLINT" json:"year" example:"1999"`
	// длина фильма
	FilmLength	string		`gorm:"null;type:VARCHAR(10)" json:"filmLength,omitempty" example:"2:16"`
	// описание фильма
	Description	string		`gorm:"null;type:STRING" json:"description,omitempty" example:"Жизнь Томаса Андерсона разделена на две части: днём он — самый обычный офисный работник, получающий нагоняи от начальства, а ночью превращается в хакера по имени Нео, и нет места в сети, куда он бы не смог проникнуть. Но однажды всё меняется. Томас узнаёт ужасающую правду о реальности."`
	// тип фильма (фильм, сериал, видео, мини-сериал)
	Type		string		`gorm:"null;type:VARCHAR(20)" json:"type" example:"фильм"`
	// персонал фильма
	Personal	*FilmStaff 	`gorm:"not null;type:JSONB" json:"personal,omitempty"`
	// дата последнего обновления информации от Kinopoisk API (в формате RFC3339)
	UpdatedAt	time.Time	`gorm:"not null;type:TIMESTAMP;index:idx_movies_updated_at,sort:desc" json:"updatedAt" example:"2025-02-28T22:00:05.225526936Z"`

	// жанры фильма
	Genres		[]Genre		`gorm:"foreignKey:MovieID" json:"genres"`
	// юзеры, которые добавили этот фильм себе
	UserMovies	[]UserMovie	`gorm:"foreignKey:MovieID" json:"-"`
}
func (Movie) TableName() string {
  return "movies"
}


// структура для парсинга ответа от API
//easyjson:json
type RawFilmInfo struct {
	KinopoiskID		int 	`json:"kinopoiskId"`
	Title			string 	`json:"nameRu"`
	PosterURL 		string 	`json:"posterUrlPreview"`
	WebURL 			string 	`json:"webUrl"`
	RatingKinopoisk	float64 `json:"ratingKinopoisk"`
	Year			int 	`json:"year"`
	FilmLenMinutes	int 	`json:"filmLength"`
	Description		string 	`json:"description"`
	Type			string 	`json:"type"`
	Genres 			[]Genre `json:"genres"`
}
