package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/user"
	"Filmer/server/internal/app/user/repository"
	"Filmer/server/internal/app/user/usecase"
	"Filmer/server/internal/pkg/token"
	"Filmer/server/internal/pkg/validator"
)

type UserHandlerManager struct {
	validator validator.Validator
	userUC    user.Usecase
}

func NewUserHandlerManager(cfg *config.Config, dbClient *gorm.DB,
	validator validator.Validator) *UserHandlerManager {

	userRepo := repository.NewDBRepo(dbClient)
	userUC := usecase.NewUsecase(cfg, userRepo)

	return &UserHandlerManager{
		validator: validator,
		userUC:    userUC,
	}
}

// @summary		Смена пароля юзера
// @description	Установка нового пароля юзеру с подтверждением через старый пароль
// @router			/user/change-password [post]
// @id				user-change-password
// @tags			user
// @param			changePasswordIn	body	changePasswordIn	true	"changePasswordIn"
// @security		JWT
// @success		204	"No Content"
// @failure		400	"Неверный пароль"
// @failure		401	"Пустой или неправильный токен"
// @failure		403	"Истекший или невалидный токен"
// @failure		404	"Текущий юзер не найден"
func (h UserHandlerManager) ChangePassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(changePasswordIn)
		user := new(entity.User)

		// parse JSON-body
		if err = ctx.BodyParser(dataIn); err != nil {
			return fmt.Errorf("change password: %w", err)
		}
		// validate parsed data
		if err = h.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("change password: %w", err)
		}

		// get user ID from token
		user.ID, err = token.ParseUserIDFromContext(ctx)
		if err != nil {
			return fmt.Errorf("change password: %w", err)
		}
		user.Password = []byte(dataIn.CurrentPassword)

		// change user password
		if err := h.userUC.ChangePassword(user, []byte(dataIn.NewPassword)); err != nil {
			return err
		}
		return ctx.Status(http.StatusNoContent).Send(nil)
	}
}

// @summary		Получение активности юзеров
// @description	Получение для каждого юзера количества фильмов в категориях "избранное", "хочу посмотреть" и "поcмотрел"
// @router			/user/all/activity [get]
// @id				user-all-activity
// @tags			user
// @success		200	{object}	entity.UsersActivity
func (h UserHandlerManager) GetActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// get users activity
		activity, err := h.userUC.GetActivity()
		if err != nil {
			return err
		}
		return ctx.JSON(activity)
	}
}
