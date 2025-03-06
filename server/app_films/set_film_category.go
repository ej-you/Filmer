package app_films

import (
	"fmt"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"

	coreValidator "server/core/validator"
	"server/core/services"
	"server/db/schemas"
	"server/db"
)


const (
	staredCategory = "stared"
	statusCategory = "status"
)

// структура обновления категории фильма юзера
type filmCategoryUpdater struct {
	Field string
	Value any
}


//easyjson:json
type SetFilmCategoryIn struct {
	// ID фильма
	MovieID uuid.UUID `params:"movieID" validate:"required" example:"86ae41a4-612a-4157-ba82-405872d1d264"`
}


// изменение категории фильма юзера
func changeCategory(ctx *fiber.Ctx, updater filmCategoryUpdater) error {
	var err error
	var dataIn SetFilmCategoryIn
	var userMovie schemas.UserMovie

	// парсинг path-параметров
	if err = ctx.ParamsParser(&dataIn); err != nil {
		return fmt.Errorf("change film category to %v: %w", updater, err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("change film category to %v: %w", updater, err)
	}

	dbConn := db.GetConn()
	// заполнение структуры фильма юзера данными для выборки по первичному ключу
	userMovie.UserID = services.ParseUserIDFromContext(ctx)
	userMovie.MovieID = dataIn.MovieID

	// проверка того, что фильм с переданным ID есть в БД
	var foundMovie int64
	selectCountResult := dbConn.Table("movies").Where("id = ?", dataIn.MovieID).Count(&foundMovie)
	if err = selectCountResult.Error; err != nil {
		return fmt.Errorf("change film category to %v: %w", updater, fiber.NewError(500, "failed to find movie with given id: " + err.Error()))
	}
	// если фильм не найден
	if foundMovie == 0 {
		return fmt.Errorf("change film category to %v: %w", updater, fiber.NewError(404, "movie with given id was not found"))
	}

	// получение записи или создание, если такой ещё нет
	selectResult := dbConn.Where(userMovie).FirstOrCreate(&userMovie)
	if err = selectResult.Error; err != nil {
		return fmt.Errorf("change film category to %v: %w", updater, fiber.NewError(500, "failed to find or create user movie: " + err.Error()))
	}

	if updater.Field == staredCategory {
		// если фильм уже в избранном/нет
		if userMovie.Stared == updater.Value {
			return ctx.Status(200).JSON(userMovie)
		}
	} else if updater.Field == statusCategory {
		// если у фильма уже статус updater.Value
		if userMovie.Status == updater.Value {
			return ctx.Status(200).JSON(userMovie)
		}
	}

	// изменяем категорию у фильма
	updateResult := dbConn.Model(&userMovie).Update(updater.Field, updater.Value)
	if err = updateResult.Error; err != nil {
		return fmt.Errorf("change film category to %v: %w", updater, fiber.NewError(500, "failed to update user movie: " + err.Error()))
	}
	return ctx.Status(200).JSON(userMovie)
}


//	@summary		Добавление фильма юзера в избранное
//	@description	Добавление фильма юзера в избранное по ID фильма
//	@router			/films/{movieID}/star [post]
//	@id				films-set-star
//	@tags			films
//	@security		JWT
//	@param			movieID	path		string	true	"ID фильма"
//	@success		200		{object}	schemas.UserMovie
//	@failure		401		"Пустой или неправильный токен"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильм не найден"
func Star(ctx *fiber.Ctx) error {
	return changeCategory(ctx, filmCategoryUpdater{staredCategory, true})
}

//	@summary		Удаление фильма юзера из избранного
//	@description	Удаление фильма юзера из избранного по ID фильма
//	@router			/films/{movieID}/unstar [post]
//	@id				films-set-unstar
//	@tags			films
//	@security		JWT
//	@param			movieID	path		string	true	"ID фильма"
//	@success		200		{object}	schemas.UserMovie
//	@failure		401		"Пустой или неправильный токен"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильм не найден"
func Unstar(ctx *fiber.Ctx) error {
	return changeCategory(ctx, filmCategoryUpdater{staredCategory, false})
}

//	@summary		Удаление фильма юзера из списков "хочу посмотреть" и "посмотрел"
//	@description	Удаление фильма юзера из списков "хочу посмотреть" и "посмотрел" по ID фильма
//	@router			/films/{movieID}/clear [post]
//	@id				films-set-clear
//	@tags			films
//	@security		JWT
//	@param			movieID	path		string	true	"ID фильма"
//	@success		200		{object}	schemas.UserMovie
//	@failure		401		"Пустой или неправильный токен"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильм не найден"
func Clear(ctx *fiber.Ctx) error {
	return changeCategory(ctx, filmCategoryUpdater{statusCategory, int8(0)})
}

//	@summary		Добавление фильма юзера в список "хочу посмотреть"
//	@description	Добавление фильма юзера в список "хочу посмотреть" по ID фильма
//	@router			/films/{movieID}/want [post]
//	@id				films-set-want
//	@tags			films
//	@security		JWT
//	@param			movieID	path		string	true	"ID фильма"
//	@success		200		{object}	schemas.UserMovie
//	@failure		401		"Пустой или неправильный токен"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильм не найден"
func SetWant(ctx *fiber.Ctx) error {
	return changeCategory(ctx, filmCategoryUpdater{statusCategory, int8(1)})
}

//	@summary		Добавление фильма юзера в список "посмотрел"
//	@description	Добавление фильма юзера в список "посмотрел" по ID фильма
//	@router			/films/{movieID}/watched [post]
//	@id				films-set-watched
//	@tags			films
//	@security		JWT
//	@param			movieID	path		string	true	"ID фильма"
//	@success		200		{object}	schemas.UserMovie
//	@failure		401		"Пустой или неправильный токен"
//	@failure		403		"Истекший или невалидный токен"
//	@failure		404		"Фильм не найден"
func SetWatched(ctx *fiber.Ctx) error {
	return changeCategory(ctx, filmCategoryUpdater{statusCategory, int8(2)})
}
