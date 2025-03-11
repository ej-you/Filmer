package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/utils"
	"Filmer/server/pkg/validator"

	"Filmer/server/internal/auth"
	"Filmer/server/internal/auth/repository"
	"Filmer/server/internal/auth/usecase"
)

// Auth handlers manager
type AuthHandlerManager struct {
	validator validator.Validator
	authUC    auth.Usecase
}

// AuthHandlerManager constructor
func NewAuthHandlerManager(cfg *config.Config, dbClient *gorm.DB, cacheClient cache.Cache,
	validator validator.Validator) *AuthHandlerManager {

	authRepo := repository.NewRepository(dbClient)
	authCacheRepo := repository.NewCacheRepository(cfg, cacheClient)
	authUC := usecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &AuthHandlerManager{
		validator: validator,
		authUC:    authUC,
	}
}

// @summary		Регистрация юзера
// @description	Регистрация нового юзера с почтой и паролем
// @router			/user/sign-up [post]
// @id				user-sign-up
// @tags			user
// @accept			json
// @produce		json
// @param			authIn	body		authIn	true	"authIn"
// @success		201		{object}	entity.UserWithToken
// @failure		409		"Юзер с введенной почтой уже зарегистрирован"
func (ahm AuthHandlerManager) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(authIn)
		user := new(entity.User)

		// parse JSON-body
		if err = ctx.BodyParser(dataIn); err != nil {
			return fmt.Errorf("sign up: %w", err)
		}
		// validate parsed data
		if err = ahm.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("sign up: %w", err)
		}

		user.Email = dataIn.Email
		user.Password = []byte(dataIn.Password)
		// sign up new user
		userWithToken, err := ahm.authUC.SignUp(user)
		if err != nil {
			return err
		}
		return ctx.Status(http.StatusCreated).JSON(userWithToken)
	}
}

// @summary		Вход для юзера
// @description	Вход для существующего юзера по почте и паролю
// @router			/user/login [post]
// @id				user-login
// @tags			user
// @accept			json
// @produce		json
// @param			authIn	body		authIn	true	"authIn"
// @success		200		{object}	entity.UserWithToken
// @failure		401		"Неверный пароль для учетной записи юзера"
// @failure		404		"Юзер с введенной почтой не найден"
func (ahm AuthHandlerManager) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(authIn)
		user := new(entity.User)

		// parse JSON-body
		if err = ctx.BodyParser(dataIn); err != nil {
			return fmt.Errorf("login user: %w", err)
		}
		// validate parsed data
		if err = ahm.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("login user: %w", err)
		}

		user.Email = dataIn.Email
		user.Password = []byte(dataIn.Password)
		// log in existing user
		userWithToken, err := ahm.authUC.Login(user)
		if err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(userWithToken)
	}
}

// @summary		Выход юзера
// @description	Выход юзера (помещение JWT-token'а текущей сессии юзера в черный список)
// @router			/user/logout [post]
// @id				user-logout
// @tags			user
// @security		JWT
// @success		204	"No Content"
// @failure		401	"Пустой или неправильный токен"
// @failure		403	"Истекший или невалидный токен"
func (ahm AuthHandlerManager) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse access token
		token := utils.ParseRawTokenFromContext(ctx)

		// put token to blacklist
		if err := ahm.authUC.Logout(token); err != nil {
			return err
		}
		return ctx.Status(http.StatusNoContent).Send(nil)
	}
}
