// Package usecase contains kinopoisk usecase implementation.
package usecase

import (
	"fmt"

	"Filmer/server/internal/app/kinopoisk"
)

const (
	_officialAPIName   = "official"
	_unofficialAPIName = "unofficial"
)

var _ kinopoisk.Usecase = (*usecase)(nil)

// kinopoisk.Usecase implementation.
type usecase struct {
	kinopoiskCacheRepo kinopoisk.CacheRepo
}

// Returns kinopoisk.Usecase interface.
func NewUsecase(kinopoiskCacheRepo kinopoisk.CacheRepo) kinopoisk.Usecase {
	return &usecase{
		kinopoiskCacheRepo: kinopoiskCacheRepo,
	}
}

// Set api limit for official API.
func (u usecase) SetOfficialAPILimit() error {
	if err := u.kinopoiskCacheRepo.SetAPILimit(_officialAPIName); err != nil {
		return fmt.Errorf("official api: %w", err)
	}
	return nil
}

// Returns true, if official API limit is reached.
func (u usecase) IsOfficialAPILimitReached() (bool, error) {
	reached, err := u.kinopoiskCacheRepo.IsAPILimitReached(_officialAPIName)
	if err != nil {
		return false, fmt.Errorf("official api: %w", err)
	}
	return reached, nil
}

// Set api limit for unofficial API.
func (u usecase) SetUnofficialAPILimit() error {
	if err := u.kinopoiskCacheRepo.SetAPILimit(_unofficialAPIName); err != nil {
		return fmt.Errorf("unofficial api: %w", err)
	}
	return nil
}

// Returns true, if unofficial API limit is reached.
func (u usecase) IsUnofficialAPILimitReached() (bool, error) {
	reached, err := u.kinopoiskCacheRepo.IsAPILimitReached(_unofficialAPIName)
	if err != nil {
		return false, fmt.Errorf("unofficial api: %w", err)
	}
	return reached, nil
}
