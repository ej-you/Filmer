package http_errors

import (
	fiber "github.com/gofiber/fiber/v2"
)


// интерфейс http ошибки
type HTTPError interface {
	Error() string
	StatusCode() int
}


// структура http ошибки
type HTTPErr struct {
	fiberError *fiber.Error
}

// конструктор для типа интерфейса HTTPError
func NewHTTPError(statusCode int, message string) HTTPError {
	return &HTTPErr{
		fiberError: fiber.NewError(statusCode, message),
	}
}

func (this HTTPErr) StatusCode() int {
	return this.fiberError.Code
}

func (this HTTPErr) Error() string {
	return this.fiberError.Message
}
