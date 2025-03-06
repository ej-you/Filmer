package app_user

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"server/core/services"
	"server/redis"
)


//	@summary		Выход юзера
//	@description	Выход юзера (помещение JWT-token'а текущей сессии юзера в черный список)
//	@router			/user/logout [post]
//	@id				user-logout
//	@tags			user
//	@security		JWT
//	@success		204	"No Content"
//	@failure		401	"Пустой или неправильный токен"
//	@failure		403	"Истекший или невалидный токен"
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
