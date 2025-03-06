package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/validator"
	"Filmer/server/pkg/utils"
	
	"Filmer/server/internal/auth"
)


//	@summary		Регистрация юзера
//	@description	Регистрация нового юзера с почтой и паролем
//	@router			/user/sign-up [post]
//	@id				user-sign-up
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			SignUpIn	body		SignUpIn	true	"SignUpIn"
//	@success		201			{object}	schemas.User
//	@failure		409			"Юзер с введенной почтой уже зарегистрирован"
func SignUp(authUC auth.Usecase, valid validator.Validator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(signUpIn)
		user := new(entity.User)
		dataOut := new(entity.UserWithToken)

		// парсинг JSON-body
		if err = ctx.BodyParser(dataIn); err != nil {
			return fmt.Errorf("sign up: %w", err)
		}
		// валидация полученной структуры
		if err = valid.Validate(dataIn); err != nil {
			return fmt.Errorf("sign up: %w", err)
		}

		user.Email = dataIn.Email
		user.Password = []byte(dataIn.Password)
		// регистрация нового юзера
		dataOut, err = authUC.SignUp(user)
		if err != nil {
			return err
		}
		return ctx.Status(201).JSON(dataOut)
	}
}

//	@summary		Вход для юзера
//	@description	Вход для существующего юзера по почте и паролю
//	@router			/user/login [post]
//	@id				user-login
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			LoginIn	body		LoginIn	true	"LoginIn"
//	@success		200		{object}	schemas.User
//	@failure		401		"Неверный пароль для учетной записи юзера"
//	@failure		404		"Юзер с введенной почтой не найден"
func Login(authUC auth.Usecase, valid validator.Validator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(loginIn)
		user := new(entity.User)
		dataOut := new(entity.UserWithToken)

		// парсинг JSON-body
		if err = ctx.BodyParser(dataIn); err != nil {
			return fmt.Errorf("login user: %w", err)
		}
		// валидация полученной структуры
		if err = valid.Validate(dataIn); err != nil {
			return fmt.Errorf("login user: %w", err)
		}

		user.Email = dataIn.Email
		user.Password = []byte(dataIn.Password)
		// вход существующего юзера
		dataOut, err = authUC.Login(user)
		if err != nil {
			return err
		}
		return ctx.Status(200).JSON(dataOut)
	}
}

//	@summary		Выход юзера
//	@description	Выход юзера (помещение JWT-token'а текущей сессии юзера в черный список)
//	@router			/user/logout [post]
//	@id				user-logout
//	@tags			user
//	@security		JWT
//	@success		204	"No Content"
//	@failure		401	"Пустой или неправильный токен"
//	@failure		403	"Истекший или невалидный токен"
func Logout(authUC auth.Usecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// получение access-токена
		token := utils.ParseRawTokenFromContext(ctx)

		// помещение токена в чёрный список
		if err := authUC.Logout(token); err != nil {
			return err
		}
		return ctx.Status(204).Send(nil)
	}
}
