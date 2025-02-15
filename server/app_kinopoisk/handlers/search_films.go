package handlers

import (
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
)


func SearchFilms(ctx *fiber.Ctx) error {
	return ctx.JSON(kinopoiskAPI.Person{
		ID: 12345,
		Name: "Quentin Tarantino",
		ImgUrl: "not-url",
	})
}
