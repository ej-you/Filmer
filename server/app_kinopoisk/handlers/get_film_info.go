package handlers

import (
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
	"server/core/services"
)

// var filmTypes = map[string]string {
// 	"FILM": "фильм",
// 	"TV_SERIES": "сериал",
// 	"VIDEO": "видео",
// 	"MINI_SERIES": "мини-сериал",
// 	"TV_SHOW": "сериал",
// }


func GetFilmInfo(ctx *fiber.Ctx) error {
	return ctx.JSON(kinopoiskAPI.Person{
		ID: 12345,
		Name: "Quentin Tarantino",
		ImgUrl: services.ParseUserIDFromContext(ctx).String(),
	})
}
