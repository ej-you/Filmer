// Package usecase contains personal usecase implementation.
package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/kinopoisk"
	"Filmer/server/internal/app/staff"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/utils"
)

var _ staff.Usecase = (*usecase)(nil)

// staff.Usecase implementation.
type usecase struct {
	logger             logger.Logger
	staffKinopoiskRepo staff.KinopoiskRepo
	kinopoiskUC        kinopoisk.Usecase
}

func New(logger logger.Logger, staffKinopoiskRepo staff.KinopoiskRepo,
	kinopoiskUC kinopoisk.Usecase) staff.Usecase {

	return &usecase{
		logger:             logger,
		staffKinopoiskRepo: staffKinopoiskRepo,
		kinopoiskUC:        kinopoiskUC,
	}
}

// Get full person info by its ID.
// Person ID must be presented.
// Fill given person struct.
func (u *usecase) GetByID(person *entity.PersonFull) error {
	// check API limit is reached
	isLimitReached, err := u.kinopoiskUC.IsUnofficialAPILimitReached()
	if err != nil {
		return fmt.Errorf("staff usecase.GetByID: %w", err)
	}
	// if daily API limit is reached
	if isLimitReached {
		return httperror.New(http.StatusPaymentRequired, "daily API limit is reached",
			errors.New("staff usecase.GetByID: api limit"),
		)
	}
	// get person info from API
	if err := u.staffKinopoiskRepo.GetFullInfoByID(person); err != nil {
		// set API limit to cache if gotten error is 402 error
		return fmt.Errorf("staff usecase.GetByID: %w",
			utils.HandleUnofficialAPIError(err, u.kinopoiskUC))
	}
	return nil
}
