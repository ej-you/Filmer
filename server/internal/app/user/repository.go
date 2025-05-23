package user

import (
	"Filmer/server/internal/app/entity"
)

type DBRepo interface {
	GetUserByID(user *entity.User) error
	UpdateUser(user *entity.User) error
}
