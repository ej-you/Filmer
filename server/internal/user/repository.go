package user

import (
	"Filmer/server/internal/entity"
)

type Repository interface {
	GetUserByID(user *entity.User) error
	UpdateUser(user *entity.User) error
}
