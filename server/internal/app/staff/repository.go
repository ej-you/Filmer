package staff

import (
	"Filmer/server/internal/app/entity"
)

type KinopoiskRepo interface {
	GetFullInfoByID(person *entity.PersonFull) error
}
