package utils

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"

	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/logger"
)


// структура для возврата ошибки клиенту
//easyjson:json
type errorResponse struct {
	StatusCode	int `json:"-"`
	Message 	string `json:"message"`
}


// кастомный обработчик ошибок
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.NewLogger().Error(err)

	var errResp errorResponse

	var fiberErr *fiber.Error
	var httpErr httpError.HTTPError

	switch {
		// если ошибка *fiber.Error
		case errors.As(err, &fiberErr):
			errResp.StatusCode = fiberErr.Code
			errResp.Message = fiberErr.Message
		// если ошибка httpError.HTTPError
		case errors.As(err, &httpErr):
			errResp.StatusCode = httpErr.StatusCode()
			errResp.Message = httpErr.Error()
		default:
			errResp.StatusCode = 500
			errResp.Message = err.Error()
	}
	// отправка структуры ошибки
	return ctx.Status(errResp.StatusCode).JSON(errResp)
}
