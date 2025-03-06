package entity

import (
	"time"

	"github.com/google/uuid"
)


//easyjson:json
// @description модель юзера
type User struct {
	// uuid юзера
	ID			uuid.UUID 	`gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	// почта юзера
	Email		string 		`gorm:"not null;type:VARCHAR(100);uniqueIndex:uni_users_email" json:"email" example:"user@gmail.com"`
	// хэш пароля юзера
	Password	[]byte 		`gorm:"not null;type:BYTES" json:"-"`
	// дата создания аккаунта
	CreatedAt	time.Time 	`gorm:"not null;type:TIMESTAMP" json:"-"`
	
	// фильмы юзера
	UserMovies	UserMovies	`gorm:"foreignKey:UserID" json:"-"`
}
func (User) TableName() string {
  return "users"
}

//easyjson:json
// @description получаемые данные о юзере с токеном
type UserWithToken struct {
	// данные о юзере
	User		*User `json:"user"`
	// access токен, генерируемый после успешной аутентификации
	AccessToken	string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3NjA0MzIsInVzZXJJRCI6IjU4MzU0ZGJhLWUyY2MtNDI4My04MjAxLWNjZGRiZTkzOGExNSJ9.VwA9d9lao0Xgl5c3H9VNM8JVtUKDs47YEItb6MZlkCw"`
}