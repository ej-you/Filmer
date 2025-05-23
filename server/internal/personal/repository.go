package personal

import (
	"Filmer/server/internal/entity"
)

type CacheRepository interface {
	SetPersonInfo(person *entity.PersonFull) error
	GetPersonInfo(person *entity.PersonFull) (bool, error)
}

type KinopoiskWebAPIRepository interface {
	GetFullInfoByID(person *entity.PersonFull) error
}
