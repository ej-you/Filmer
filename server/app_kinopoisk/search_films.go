package app_kinopoisk

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
	coreValidator "server/core/validator"
)


//easyjson:json
// данные для поиска фильмов по ключевому слову
type SearchFilmsIn struct {
	// ключевое слово для поиска
	Query 	string `query:"q" validate:"required" example:"матрица"`
	// страница
	Page 	int `query:"page" validate:"required,min=1" example:"1"`
}


//	@summary		Поиск фильмов
//	@description	Поиск фильмов по ключевому слову с пагинацией
//	@router			/kinopoisk/films/search [get]
//	@id				kinopoisk-search-films
//	@tags			kinopoisk-films
//	@security		JWT
//	@param			q		query		string	true	"ключевое слово (Например: матрица)"
//	@param			page	query		int		true	"страница поиска (Например: 1)"
//	@success		200		{object}	kinopoiskAPI.SearchedFilms
//	@failure		401		"Пустой или неправильный токен"
//	@failure		402		"Превышен дневной лимит запросов к Kinopoisk API"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильмы не найдены"
//	@failure		429		"Слишком много запросов. Лимит 5 запросов в секунду"
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
	// получение данных от API
	if err = kinopoiskAPI.SearchFilmsByKeyword(dataIn.Query, dataIn.Page, &searchedFilms); err != nil {
		return fmt.Errorf("search films: %w", err)
	}
	return ctx.Status(200).JSON(searchedFilms)
}
