package schemas

import (
	"time"

	"github.com/google/uuid"
)


// модель юзера
//easyjson:json
// @description получаемые данные о юзере
type User struct {
	// uuid юзера
	ID			uuid.UUID 	`gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	// почта юзера
	Email		string 		`gorm:"not null;type:VARCHAR(100);unique" json:"email" example:"user@gmail.com"`
	// хэш пароля юзера
	Password	[]byte 		`gorm:"not null;type:BYTES" json:"-"`
	// дата создания аккаунта
	CreatedAt	time.Time 	`gorm:"not null;type:TIMESTAMP" json:"-"`
	// access токен, генерируемый после успешной аутентификации
	AccessToken	string 		`gorm:"-" json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3NjA0MzIsInVzZXJJRCI6IjU4MzU0ZGJhLWUyY2MtNDI4My04MjAxLWNjZGRiZTkzOGExNSJ9.VwA9d9lao0Xgl5c3H9VNM8JVtUKDs47YEItb6MZlkCw"`
	
	// фильмы юзера
	UserMovies	UserMovies	`gorm:"foreignKey:UserID" json:"-"`
}
