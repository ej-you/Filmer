package user

import (
	"Filmer/server/internal/entity"
)

type Usecase interface {
	ChangePassword(user *entity.User, newPassword []byte) error
}
