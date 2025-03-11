package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/utils"
)

const (
	starConst    = true
	unstarConst  = false
	clearConst   = 0
	wantConst    = 1
	watchedConst = 2
)

// @summary		Добавление фильма юзера в избранное
// @description	Добавление фильма юзера в избранное по ID фильма
// @router			/films/{movieID}/star [post]
// @id				films-set-star
// @tags			user-movie
// @security		JWT
// @param			movieID	path		string	true	"ID фильма"
// @success		200		{object}	entity.UserMovie
// @failure		401		"Пустой или неправильный токен"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильм не найден"
func (umhm UserMovieHandlerManager) Star() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return umhm.changeMovieStared(ctx, starConst)
	}
}

// @summary		Удаление фильма юзера из избранного
// @description	Удаление фильма юзера из избранного по ID фильма
// @router			/films/{movieID}/unstar [post]
// @id				films-set-unstar
// @tags			user-movie
// @security		JWT
// @param			movieID	path		string	true	"ID фильма"
// @success		200		{object}	entity.UserMovie
// @failure		401		"Пустой или неправильный токен"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильм не найден"
func (umhm UserMovieHandlerManager) Unstar() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return umhm.changeMovieStared(ctx, unstarConst)
	}
}

// @summary		Удаление фильма юзера из списков "хочу посмотреть" и "посмотрел"
// @description	Удаление фильма юзера из списков "хочу посмотреть" и "посмотрел" по ID фильма
// @router			/films/{movieID}/clear [post]
// @id				films-set-clear
// @tags			user-movie
// @security		JWT
// @param			movieID	path		string	true	"ID фильма"
// @success		200		{object}	entity.UserMovie
// @failure		401		"Пустой или неправильный токен"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильм не найден"
func (umhm UserMovieHandlerManager) Clear() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return umhm.changeMovieStatus(ctx, clearConst)
	}
}

// @summary		Добавление фильма юзера в список "хочу посмотреть"
// @description	Добавление фильма юзера в список "хочу посмотреть" по ID фильма
// @router			/films/{movieID}/want [post]
// @id				films-set-want
// @tags			user-movie
// @security		JWT
// @param			movieID	path		string	true	"ID фильма"
// @success		200		{object}	entity.UserMovie
// @failure		401		"Пустой или неправильный токен"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильм не найден"
func (umhm UserMovieHandlerManager) SetWant() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return umhm.changeMovieStatus(ctx, wantConst)
	}
}

// @summary		Добавление фильма юзера в список "посмотрел"
// @description	Добавление фильма юзера в список "посмотрел" по ID фильма
// @router			/films/{movieID}/watched [post]
// @id				films-set-watched
// @tags			user-movie
// @security		JWT
// @param			movieID	path		string	true	"ID фильма"
// @success		200		{object}	entity.UserMovie
// @failure		401		"Пустой или неправильный токен"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильм не найден"
func (umhm UserMovieHandlerManager) SetWatched() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return umhm.changeMovieStatus(ctx, watchedConst)
	}
}

// change stared field value in user movie
func (umhm UserMovieHandlerManager) changeMovieStared(ctx *fiber.Ctx, newStared bool) error {
	var err error
	dataIn := new(setFilmCategoryIn)
	userMovie := new(entity.UserMovie)

	// parse path-params
	if err = ctx.ParamsParser(dataIn); err != nil {
		return fmt.Errorf("change movie stared to %v: %w", newStared, err)
	}
	// validate parsed data
	if err = umhm.validator.Validate(dataIn); err != nil {
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
	err = umhm.userMovieUC.UpdateUserMovieStared(userMovie, newStared)
	if err != nil {
		return err
	}
	// clear substruct value (movie info)
	userMovie.Movie = nil
	return ctx.Status(http.StatusOK).JSON(userMovie)
}

// change status field value in user movie
func (umhm UserMovieHandlerManager) changeMovieStatus(ctx *fiber.Ctx, newStatus int8) error {
	var err error
	dataIn := new(setFilmCategoryIn)
	userMovie := new(entity.UserMovie)

	// parse path-params
	if err = ctx.ParamsParser(dataIn); err != nil {
		return fmt.Errorf("change movie status to %v: %w", newStatus, err)
	}
	// validate parsed data
	if err = umhm.validator.Validate(dataIn); err != nil {
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
	err = umhm.userMovieUC.UpdateUserMovieStatus(userMovie, newStatus)
	if err != nil {
		return err
	}
	// clear substruct value (movie info)
	userMovie.Movie = nil
	return ctx.Status(http.StatusOK).JSON(userMovie)
}
