package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/auth"
	"Filmer/server/internal/app/auth/repository"
	"Filmer/server/internal/app/auth/usecase"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/token"
	"Filmer/server/internal/pkg/validator"
)

// Auth handlers manager
type AuthHandlerManager struct {
	validator validator.Validator
	authUC    auth.Usecase
}

func NewAuthHandlerManager(cfg *config.Config, dbClient *gorm.DB, cacheStorage cache.Storage,
	validator validator.Validator) *AuthHandlerManager {

	authRepo := repository.NewDBRepo(dbClient)
	authCacheRepo := repository.NewCacheRepository(cfg, cacheStorage)
	authUC := usecase.NewUsecase(cfg, authRepo, authCacheRepo)

	return &AuthHandlerManager{
		validator: validator,
		authUC:    authUC,
	}
}

// @summary		Регистрация юзера
// @description	Регистрация нового юзера с почтой и паролем
// @router			/auth/sign-up [post]
// @id				auth-sign-up
// @tags			auth
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
// @router			/auth/login [post]
// @id				auth-login
// @tags			auth
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
// @router			/auth/logout [post]
// @id				auth-logout
// @tags			auth
// @security		JWT
// @success		204	"No Content"
// @failure		401	"Пустой или неправильный токен"
// @failure		403	"Истекший или невалидный токен"
func (ahm AuthHandlerManager) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse access token
		token := token.ParseRawTokenFromContext(ctx)

		// put token to blacklist
		if err := ahm.authUC.Logout(token); err != nil {
			return err
		}
		return ctx.Status(http.StatusNoContent).Send(nil)
	}
}
