package handlers

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


//easyjson:json
type SignUpIn struct {
	// почта юзера
	Email 	string `json:"email" validate:"required,max=100" example:"vasya2007@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=50" example:"qwerty123" minLength:"8" maxLength:"50"`
}


func SignUp(ctx *fiber.Ctx) error {
	var err error
	var dataIn SignUpIn
	var user schemas.User

	// парсинг JSON-body
	if err = ctx.BodyParser(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}

	// хэширование пароля и заполнение структуры
	passwordHash, err := services.EncodePassword(dataIn.Password)
	if err != nil {
		return err
	}
	user.Email = dataIn.Email
	user.Password = passwordHash

	// создание нового юзера в БД
	createResult := db.GetConn().Create(&user)
	if err = createResult.Error; err != nil {
		// если юзер с таким юзернеймом уже есть
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return fiber.NewError(409, "user with such email already exists")
		}
		return fiber.NewError(500, "failed to create user: " + err.Error())
	}
	// генерация access токена
	user.AccessToken, err = services.ObtainToken(user.ID)
	if err != nil {
		return err
	}
	return ctx.Status(201).JSON(user)
}
