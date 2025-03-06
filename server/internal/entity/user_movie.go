package entity

import (
	"time"

	"github.com/google/uuid"
)


// модель фильмов юзеров
//easyjson:json
// @description получаемые данные о фильме юзера
type UserMovie struct {
	// uuid юзера
	UserID	uuid.UUID 	`gorm:"not null;type:uuid;primaryKey" json:"-"`
	// uuid фильма
	MovieID	uuid.UUID 	`gorm:"not null;type:uuid;primaryKey" json:"-"`
	// в каком списке фильм (0 - ничего, 1 - хочу посмотреть, 2 - посмотрел)
	Status	int8		`gorm:"not null;type:SMALLINT" json:"status" example:"1"`
	// в избранном ли фильм
	Stared	bool		`gorm:"not null;type:BOOLEAN" json:"stared" example:"true"`
	// дата обновления записи
	UpdatedAt	time.Time	`gorm:"not null;type:TIMESTAMP" json:"-"`
	
	// ассоциации c юзером и фильмом
	User	User	`gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Movie	*Movie	`gorm:"not null;foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie,omitempty"`
}
func (UserMovie) TableName() string {
  return "user_movies"
}

//easyjson:json
type UserMovies []UserMovie
