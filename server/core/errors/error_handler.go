package errors

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"

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
		default:
			errResp.Type = "unknown"
			errResp.StatusCode = 500
			errResp.Message = err.Error()
	}

	// отправка структуры ошибки
	return ctx.Status(errResp.StatusCode).JSON(errResp)
}
