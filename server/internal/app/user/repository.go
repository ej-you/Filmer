package user

import (
	"Filmer/server/internal/app/entity"
)

type Repository interface {
	GetUserByID(user *entity.User) error
	UpdateUser(user *entity.User) error
}
