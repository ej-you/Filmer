package schemas

import (
	"time"

	"github.com/google/uuid"
)


// модель юзера
//easyjson:json
type User struct {
	// uuid юзера
	ID			uuid.UUID 	`gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	// Почта юзера
	Email		string 		`gorm:"not null;type:VARCHAR(100);unique" json:"email" example:"vasya2007@gmail.com"`
	// хэш пароля юзера
	Password	[]byte 		`gorm:"not null;type:BYTES" json:"-"`
	// дата создания аккаунта
	CreatedAt	time.Time 	`gorm:"not null" json:"createdAt"`
	// access токен, генерируемый после успешной аутентификации
	AccessToken	string 		`gorm:"-" json:"accessToken"`
	// ассоциация юзера с чатами, в которых он состоит
	// Chats 		[]Chat 	`gorm:"many2many:chat_participants" json:"-"`
}
