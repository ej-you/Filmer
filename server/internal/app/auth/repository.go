package auth

import (
	"Filmer/server/internal/app/entity"
)

type DBRepo interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(user *entity.User) error
}

type CacheRepo interface {
	SetTokenToBlacklist(token string) error
	TokenIsBlacklisted(token string) (bool, error)
}
