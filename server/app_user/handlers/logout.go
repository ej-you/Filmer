package handlers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"server/core/services"
	"server/redis"
)


// выход юзера
func Logout(ctx *fiber.Ctx) error {
	// получение access-токена
	token := services.ParseRawTokenFromContext(ctx)

	// помещение токена в чёрный список
	err := redis.SetBlacklistedToken(token)
	if err != nil {
		return fmt.Errorf("logout user: %w", err)
	}
	return ctx.Status(204).Send(nil)
}
