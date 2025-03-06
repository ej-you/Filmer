package services

import (
	"github.com/mailru/easyjson"
	fiber "github.com/gofiber/fiber/v2"
)


// для сериализации JSON с помощью easyjson
func EasyjsonEncoder(v interface{}) ([]byte, error) {
	if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}
	return nil, fiber.NewError(500, "result type does not implement easyjson.Marshaler")
}

// для десериализации JSON с помощью easyjson
func EasyjsonDecoder(data []byte, v interface{}) error {
	if um, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, um)
	}
	return fiber.NewError(500, "result type does not implement easyjson.Unmarshaler")
}
