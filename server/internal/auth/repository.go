package auth

import (
	"Filmer/server/internal/entity"
)


type Repository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindUserByEmail(user *entity.User) (*entity.User, error)
}


type CacheRepository interface {
	SetTokenToBlacklist(token string) error // cache cache.Cache, cfg *config.Config
	TokenIsBlacklisted(token string) (bool, error) // cache cache.Cache
}
