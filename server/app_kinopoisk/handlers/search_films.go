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
	Page 	int `query:"page" validate:"min=0" example:"1" minLength:"0"`
}


// поиск фильмов по ключевому слову
func SearchFilms(ctx *fiber.Ctx) error {
	var err error
	var dataIn SearchFilmsIn
	var searchedFilms kinopoiskAPI.SearchedFilms

	// парсинг query-параметров
	if err = ctx.QueryParser(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return err
	}
	// если страница не задана, то выставляем 1 по умолчанию
	if dataIn.Page == 0 {
		dataIn.Page = 1
	}

	if err = kinopoiskAPI.SearchFilmsByKeyword(dataIn.Query, dataIn.Page, &searchedFilms); err != nil {
		return err
	}
	fmt.Println("len searchedFilms:", len(searchedFilms.Films))
	return ctx.JSON(searchedFilms)
}
