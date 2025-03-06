package app_user

import (
	"fmt"
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


//easyjson:json
//	@description	данные для регистрации юзера
type SignUpIn struct {
	// почта юзера
	Email 	string `json:"email" validate:"required,max=100" example:"user@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=50" example:"qwerty123" minLength:"8" maxLength:"50"`
}


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
func SignUp(ctx *fiber.Ctx) error {
	var err error
	var dataIn SignUpIn
	var user schemas.User

	// парсинг JSON-body
	if err = ctx.BodyParser(&dataIn); err != nil {
		return fmt.Errorf("sign up user: %w", err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("sign up user: %w", err)
	}

	// хэширование пароля и заполнение структуры
	passwordHash, err := services.EncodePassword(dataIn.Password)
	if err != nil {
		return fmt.Errorf("sign up user: %w", err)
	}
	user.Email = dataIn.Email
	user.Password = passwordHash

	// создание нового юзера в БД
	createResult := db.GetConn().Create(&user)
	if err = createResult.Error; err != nil {
		// если юзер с таким юзернеймом уже есть
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return fmt.Errorf("sign up user: %w", fiber.NewError(409, "user with such email already exists"))
		}
		return fmt.Errorf("sign up user: %w", fiber.NewError(500, "failed to create user: " + err.Error()))
	}
	// генерация access токена
	user.AccessToken, err = services.ObtainToken(user.ID)
	if err != nil {
		return fmt.Errorf("sign up user %s: %w", user.Email, err)
	}
	return ctx.Status(201).JSON(user)
}
