// Package usecase contains personal usecase implementation.
package usecase

import (
	"fmt"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/staff"
	"Filmer/server/internal/pkg/logger"
)

var _ staff.Usecase = (*usecase)(nil)

// staff.Usecase implementation.
type usecase struct {
	logger             logger.Logger
	staffCacheRepo     staff.CacheRepo
	staffKinopoiskRepo staff.KinopoiskRepo
}

func New(logger logger.Logger, staffCacheRepo staff.CacheRepo,
	staffKinopoiskRepo staff.KinopoiskRepo) staff.Usecase {

	return &usecase{
		logger:             logger,
		staffCacheRepo:     staffCacheRepo,
		staffKinopoiskRepo: staffKinopoiskRepo,
	}
}

// Get full person info by its ID.
// First try to find person in cache. If there is not person in cache then
// send request to API and put gotten person to cache (with 1 day expiration).
// Person ID must be presented.
// Fill given person struct.
func (u *usecase) GetByID(person *entity.PersonFull) error {
	// get person from cache
	found, err := u.staffCacheRepo.GetPersonInfo(person)
	if found {
		return nil
	}
	if err != nil {
		u.logger.Errorf("staff usecase.GetByID: get from cache: %v", err)
	}

	if err := u.staffKinopoiskRepo.GetFullInfoByID(person); err != nil {
		return fmt.Errorf("staff usecase.GetByID: %w", err)
	}

	// save person to cache
	err = u.staffCacheRepo.SetPersonInfo(person)
	if err != nil {
		u.logger.Errorf("staff usecase.GetByID: set to cache: %v", err)
	}
	return nil
}
