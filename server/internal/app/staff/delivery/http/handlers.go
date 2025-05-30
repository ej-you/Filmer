package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/staff"
	personalRepository "Filmer/server/internal/app/staff/repository"
	personalUsecase "Filmer/server/internal/app/staff/usecase"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/validator"
)

type StaffHandlerManager struct {
	validator  validator.Validator
	personalUC staff.Usecase
}

func NewStaffHandlerManager(cfg *config.Config, jsonify jsonify.JSONify,
	logger logger.Logger, validator validator.Validator) *StaffHandlerManager {

	movieKinopoiskWebAPIRepo := personalRepository.NewKinopoiskRepo(cfg, jsonify)
	personalUC := personalUsecase.New(logger, movieKinopoiskWebAPIRepo)

	return &StaffHandlerManager{
		validator:  validator,
		personalUC: personalUC,
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
		if err := h.personalUC.GetByID(personInfo); err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(personInfo)
	}
}
