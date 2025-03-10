package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/utils"
)


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
func (this UserMovieHandlerManager) Star() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		return this.changeMovieStared(ctx, true)
	}
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
func (this UserMovieHandlerManager) Unstar() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		return this.changeMovieStared(ctx, false)
	}
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
func (this UserMovieHandlerManager) Clear() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		return this.changeMovieStatus(ctx, int8(0))
	}
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
func (this UserMovieHandlerManager) SetWant() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		return this.changeMovieStatus(ctx, int8(1))
	}
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
func (this UserMovieHandlerManager) SetWatched() fiber.Handler {
	return func (ctx *fiber.Ctx) error {
		return this.changeMovieStatus(ctx, int8(2))
	}
}


// change stared field value in user movie
func (this UserMovieHandlerManager) changeMovieStared(ctx *fiber.Ctx, newStared bool) error {
	var err error
	dataIn := new(setFilmCategoryIn)
	userMovie := new(entity.UserMovie)

	// parse path-params
	if err = ctx.ParamsParser(dataIn); err != nil {
		return fmt.Errorf("change movie stared to %v: %w", newStared, err)
	}
	// validate parsed data
	if err = this.validator.Validate(dataIn); err != nil {
		return fmt.Errorf("change movie stared to %v: %w", newStared, err)
	}
	// add necessary data to userMovie
	// parse user ID from token from ctx
	userMovie.UserID, err = utils.ParseUserIDFromContext(ctx)
	if err != nil {
		return fmt.Errorf("change movie stared to %v: %w", newStared, err)
	}
	userMovie.MovieID = dataIn.MovieID

	// change stared to newStared
	err = this.userMovieUC.UpdateUserMovieStared(userMovie, newStared)
	if err != nil {
		return err
	}
	// clear substruct value (movie info)
	userMovie.Movie = nil
	return ctx.Status(200).JSON(userMovie)
}

// change status field value in user movie
func (this UserMovieHandlerManager) changeMovieStatus(ctx *fiber.Ctx, newStatus int8) error {
	var err error
	dataIn := new(setFilmCategoryIn)
	userMovie := new(entity.UserMovie)

	// parse path-params
	if err = ctx.ParamsParser(dataIn); err != nil {
		return fmt.Errorf("change movie status to %v: %w", newStatus, err)
	}
	// validate parsed data
	if err = this.validator.Validate(dataIn); err != nil {
		return fmt.Errorf("change movie status to %v: %w", newStatus, err)
	}
	// add necessary data to userMovie
	// parse user ID from token from ctx
	userMovie.UserID, err = utils.ParseUserIDFromContext(ctx)
	if err != nil {
		return fmt.Errorf("change movie status to %v: %w", newStatus, err)
	}
	userMovie.MovieID = dataIn.MovieID

	// change status to newStatus
	err = this.userMovieUC.UpdateUserMovieStatus(userMovie, newStatus)
	if err != nil {
		return err
	}
	// clear substruct value (movie info)
	userMovie.Movie = nil
	return ctx.Status(200).JSON(userMovie)
}
