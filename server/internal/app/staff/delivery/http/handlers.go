package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	kinopoiskrepo "Filmer/server/internal/app/kinopoisk/repository"
	kinopoiskusecase "Filmer/server/internal/app/kinopoisk/usecase"
	"Filmer/server/internal/app/staff"
	personalrepo "Filmer/server/internal/app/staff/repository"
	personalusecase "Filmer/server/internal/app/staff/usecase"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/validator"
)

type StaffHandlerManager struct {
	validator validator.Validator
	staffUC   staff.Usecase
}

func NewStaffHandlerManager(cfg *config.Config, jsonify jsonify.JSONify,
	logger logger.Logger, cache cache.Storage,
	validator validator.Validator) *StaffHandlerManager {

	// init kinopoisk usecase
	kinopoiskCacheRepo := kinopoiskrepo.NewCacheRepo(cache)
	kinopoiskUC := kinopoiskusecase.NewUsecase(kinopoiskCacheRepo)
	// init staff usecase
	staffKinopoiskRepo := personalrepo.NewKinopoiskRepo(cfg, jsonify)
	staffUC := personalusecase.New(logger, staffKinopoiskRepo, kinopoiskUC)

	return &StaffHandlerManager{
		validator: validator,
		staffUC:   staffUC,
	}
}

// @summary		Получение информации о личности
// @description	Получение информации о личности по её ID
// @router			/staff/full-info/{personID} [get]
// @id				kinopoisk-get-staff-info
// @tags			staff
// @security		JWT
// @param			personID	path		int	true	"ID личности"
// @success		200			{object}	entity.PersonFull
// @failure		401			"Пустой или неправильный токен"
// @failure		402			"Превышен дневной лимит запросов к Kinopoisk API"
// @failure		403			"Истекший или невалидный токен"
// @failure		404			"Личность не найдена"
// @failure		429			"Слишком много запросов. Лимит 5 запросов в секунду"
func (h StaffHandlerManager) GetPersonInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// var err error
		dataIn := &getPersonInfoIn{}
		personInfo := &entity.PersonFull{}

		// parse path-params
		if err := ctx.ParamsParser(dataIn); err != nil {
			return fmt.Errorf("get person info: %w", err)
		}
		// validate parsed data
		if err := h.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("get person info: %w", err)
		}

		personInfo.ID = dataIn.PersonID
		// get person info (from cache or from API)
		if err := h.staffUC.GetByID(personInfo); err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(personInfo)
	}
}
