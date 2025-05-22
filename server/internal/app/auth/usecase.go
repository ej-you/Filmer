package auth

import (
	"Filmer/server/internal/app/entity"
)

type Usecase interface {
	SignUp(user *entity.User) (*entity.UserWithToken, error)
	Login(user *entity.User) (*entity.UserWithToken, error)
	Logout(token string) error
	RestrictBlacklistedToken(token string) error
}
