// Package usecase contains personal usecase implementation.
package usecase

import (
	"github.com/pkg/errors"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/staff"
	"Filmer/server/internal/pkg/logger"
)

var _ staff.Usecase = (*usecase)(nil)

// staff.Usecase implementation.
type usecase struct {
	logger             logger.Logger
	staffKinopoiskRepo staff.KinopoiskRepo
}

func New(logger logger.Logger,
	staffKinopoiskRepo staff.KinopoiskRepo) staff.Usecase {

	return &usecase{
		logger:             logger,
		staffKinopoiskRepo: staffKinopoiskRepo,
	}
}

// Get full person info by its ID.
// Person ID must be presented.
// Fill given person struct.
func (u *usecase) GetByID(person *entity.PersonFull) error {
	err := u.staffKinopoiskRepo.GetFullInfoByID(person)
	return errors.Wrap(err, "staff usecase.GetByID") // err OR nil
}
