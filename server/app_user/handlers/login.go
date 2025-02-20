package handlers

import (
	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


//easyjson:json
type LoginIn struct {
	// почта юзера
	Email 	string `json:"email" validate:"required,max=100" example:"vasya2007@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=50" example:"qwerty123" minLength:"8" maxLength:"50"`
}


func Login(ctx *fiber.Ctx) error {
	var err error
	var dataIn LoginIn
	var user schemas.User

	// парсинг JSON-body
	if err = ctx.BodyParser(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// получение юзера из БД по email
	selectResult := db.GetConn().Where("email = ?", dataIn.Email).First(&user)
	if err = selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return fiber.NewError(404, "user with such email was not found")
		}
		return fiber.NewError(500, "failed to create user: " + err.Error())
	}
	// сверка введённого пароля с хэшем из БД
	if !services.PasswordIsCorrect(dataIn.Password, user.Password) {
		return fiber.NewError(400, "invalid password")
	}
	
	// генерация access токена
	user.AccessToken, err = services.ObtainToken(user.ID)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(user)
}
