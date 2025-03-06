package app_user

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


//easyjson:json
// @description данные для входа юзера
type LoginIn struct {
	// почта юзера
	Email 	string `json:"email" validate:"required,max=100" example:"user@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=50" example:"qwerty123" minLength:"8" maxLength:"50"`
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
//	@failure		400		"Неверный пароль для учетной записи юзера"
//	@failure		404		"Юзер с введенной почтой не найдены"
func Login(ctx *fiber.Ctx) error {
	var err error
	var dataIn LoginIn
	var user schemas.User

	// парсинг JSON-body
	if err = ctx.BodyParser(&dataIn); err != nil {
		return fmt.Errorf("login user: %w", err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("login user: %w", err)
	}

	// получение юзера из БД по email
	selectResult := db.GetConn().Where("email = ?", dataIn.Email).First(&user)
	if err = selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if errors.Is(err, db.NotFoundError) {
			return fmt.Errorf("login user: %w", fiber.NewError(404, "user with such email was not found"))
		}
		return fmt.Errorf("login user: %w", fiber.NewError(500, "failed to get user: " + err.Error()))
	}
	// сверка введённого пароля с хэшем из БД
	if !services.PasswordIsCorrect(dataIn.Password, user.Password) {
		return fmt.Errorf("login user %s: %w", user.Email, fiber.NewError(400, "invalid password"))
	}
	
	// генерация access токена
	user.AccessToken, err = services.ObtainToken(user.ID)
	if err != nil {
		return fmt.Errorf("login user %s: %w", user.Email, err)
	}
	return ctx.Status(200).JSON(user)
}
