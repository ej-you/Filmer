package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	"Filmer/server/internal/personal"
	personalRepository "Filmer/server/internal/personal/repository"
	personalUsecase "Filmer/server/internal/personal/usecase"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/logger"
	"Filmer/server/pkg/validator"
)

// Personal handlers manager.
type PersonalHandlerManager struct {
	validator  validator.Validator
	personalUC personal.Usecase
}

func NewPersonalHandlerManager(cfg *config.Config, jsonify jsonify.JSONify, logger logger.Logger,
	cache cache.Cache, validator validator.Validator) *PersonalHandlerManager {

	movieCacheRepo := personalRepository.NewCacheRepository(cache, jsonify)
	movieKinopoiskWebAPIRepo := personalRepository.NewKinopoiskWebAPIRepository(cfg, jsonify)
	personalUC := personalUsecase.New(logger, movieCacheRepo, movieKinopoiskWebAPIRepo)

	return &PersonalHandlerManager{
		validator:  validator,
		personalUC: personalUC,
	}
}

// @summary		Получение информации о личности
// @description	Получение информации о личности по её ID
// @router			/personal/full-info/{personID} [get]
// @id				kinopoisk-get-person-info
// @tags			personal
// @security		JWT
// @param			personID	path		int	true	"ID личности"
// @success		200			{object}	entity.PersonFull
// @failure		401			"Пустой или неправильный токен"
// @failure		402			"Превышен дневной лимит запросов к Kinopoisk API"
// @failure		403			"Истекший или невалидный токен"
// @failure		404			"Личность не найдена"
// @failure		429			"Слишком много запросов. Лимит 5 запросов в секунду"
func (p PersonalHandlerManager) GetPersonInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// var err error
		dataIn := &getPersonInfoIn{}
		personInfo := &entity.PersonFull{}

		// parse path-params
		if err := ctx.ParamsParser(dataIn); err != nil {
			return fmt.Errorf("get person info: %w", err)
		}
		// validate parsed data
		if err := p.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("get person info: %w", err)
		}

		personInfo.ID = dataIn.PersonID
		// get person info (from cache or from API)
		if err := p.personalUC.GetByID(personInfo); err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(personInfo)
	}
}
