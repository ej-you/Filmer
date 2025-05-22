// Package usecase contains personal usecase implementation.
package usecase

import (
	"fmt"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/personal"
	"Filmer/server/internal/pkg/logger"
)

var _ personal.Usecase = (*personalUsecase)(nil)

// personal.Usecase implementation.
type personalUsecase struct {
	logger                      logger.Logger
	personalCacheRepo           personal.CacheRepository
	personalKinopoiskWebAPIRepo personal.KinopoiskWebAPIRepository
}

func New(logger logger.Logger, personalCacheRepo personal.CacheRepository,
	personalKinopoiskWebAPIRepo personal.KinopoiskWebAPIRepository) personal.Usecase {

	return &personalUsecase{
		logger:                      logger,
		personalCacheRepo:           personalCacheRepo,
		personalKinopoiskWebAPIRepo: personalKinopoiskWebAPIRepo,
	}
}

// Get full person info by its ID.
// First try to find person in cache. If there is not person in cache then
// send request to API and put gotten person to cache (with 1 day expiration).
// Person ID must be presented.
// Fill given person struct.
func (p *personalUsecase) GetByID(person *entity.PersonFull) error {
	// get person from cache
	found, err := p.personalCacheRepo.GetPersonInfo(person)
	if found {
		return nil
	}
	if err != nil {
		p.logger.Errorf("personalUsecase.GetByID: get from cache: %v", err)
	}

	if err := p.personalKinopoiskWebAPIRepo.GetFullInfoByID(person); err != nil {
		return fmt.Errorf("personalUsecase.GetByID: %w", err)
	}

	// save person to cache
	err = p.personalCacheRepo.SetPersonInfo(person)
	if err != nil {
		p.logger.Errorf("personalUsecase.GetByID: set to cache: %v", err)
	}
	return nil
}
