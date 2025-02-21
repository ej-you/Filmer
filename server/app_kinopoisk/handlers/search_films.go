package handlers

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
	coreValidator "server/core/validator"
)


//easyjson:json
type SearchFilmsIn struct {
	// слово для поиска
	Query 	string `query:"q" validate:"required" example:"дэдпул"`
	// страница
	Page 	int `query:"page" validate:"required,min=1" example:"1"`
}


// поиск фильмов по ключевому слову
func SearchFilms(ctx *fiber.Ctx) error {
	var err error
	var dataIn SearchFilmsIn
	var searchedFilms kinopoiskAPI.SearchedFilms

	// парсинг query-параметров
	if err = ctx.QueryParser(&dataIn); err != nil {
		return fmt.Errorf("search films: %w", err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("search films: %w", err)
	}

	if err = kinopoiskAPI.SearchFilmsByKeyword(dataIn.Query, dataIn.Page, &searchedFilms); err != nil {
		return fmt.Errorf("search films: %w", err)
	}
	return ctx.JSON(searchedFilms)
}
