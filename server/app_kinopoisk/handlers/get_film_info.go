package handlers

import (
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
	"server/core/services"
)


func GetFilmInfo(ctx *fiber.Ctx) error {
	return ctx.JSON(kinopoiskAPI.Person{
		ID: 12345,
		Name: "Quentin Tarantino",
		ImgUrl: services.ParseUserIDFromContext(ctx).String(),
	})
}
