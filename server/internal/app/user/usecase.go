package user

import (
	"Filmer/server/internal/app/entity"
)

type Usecase interface {
	ChangePassword(user *entity.User, newPassword []byte) error
	GetActivity() (entity.UsersActivity, error)
}
