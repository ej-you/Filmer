package staff

import (
	"Filmer/server/internal/app/entity"
)

type CacheRepo interface {
	SetPersonInfo(person *entity.PersonFull) error
	GetPersonInfo(person *entity.PersonFull) (bool, error)
}

type KinopoiskRepo interface {
	GetFullInfoByID(person *entity.PersonFull) error
}
