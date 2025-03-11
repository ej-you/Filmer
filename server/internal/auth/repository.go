package auth

import (
	"Filmer/server/internal/entity"
)

type Repository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(user *entity.User) error
}

type CacheRepository interface {
	SetTokenToBlacklist(token string) error
	TokenIsBlacklisted(token string) (bool, error)
}
