package errors

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	fiberJWT "github.com/gofiber/contrib/jwt"

	coreValidator "server/core/validator"
	"server/settings"
)


// структура для возврата ошибки клиенту
//easyjson:json
type ErrorResponse struct {
	StatusCode	int `json:"-"`
	Type 		string `json:"type"`
	Message 	string `json:"message"`
}


// кастомный обработчик ошибок
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	settings.ErrorLog.Print(err)

	var errResp ErrorResponse
	errResp.Type = "error"

	var fiberErr *fiber.Error

	switch {
		// если ошибка *fiber.Error
		case errors.As(err, &fiberErr):
			errResp.StatusCode = fiberErr.Code
			errResp.Message = fiberErr.Message
		// если ошибка истёкшего токена
		case errors.Is(err, jwt.ErrTokenExpired):
			errResp.Type = "token"
			errResp.StatusCode = 403
			errResp.Message = "token is expired"
		// если ошибка отсутствия токена
		case errors.Is(err, fiberJWT.ErrJWTMissingOrMalformed):
			errResp.Type = "token"
			errResp.StatusCode = 401
			errResp.Message = "token is missing or malformed"
		default:
			errResp.Type = "unknown"
			errResp.StatusCode = 500
			errResp.Message = err.Error()
	}

	// проверка на ошибки валидации
	errString, ok := coreValidator.GetValidator().GetStringFromValidationError(err)
	if ok {
		errResp.Type = "validateError"
		errResp.StatusCode = 400
		errResp.Message = errString
	}

	// отправка структуры ошибки
	return ctx.Status(errResp.StatusCode).JSON(errResp)
}
