package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	movieRepository "Filmer/server/internal/movie/repository"
	movieUsecase "Filmer/server/internal/movie/usecase"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/logger"
	"Filmer/server/pkg/utils"
	"Filmer/server/pkg/validator"

	userMovie "Filmer/server/internal/user_movie"
	"Filmer/server/internal/user_movie/repository"
	"Filmer/server/internal/user_movie/usecase"
)

const (
	getStaredConst  = "stared"
	getWantConst    = "want"
	getWatchedConst = "watched"
)

// User movie handlers manager
type UserMovieHandlerManager struct {
	validator   validator.Validator
	userMovieUC userMovie.Usecase
}

// UserMovieHandlerManager constructor
func NewUserMovieHandlerManager(cfg *config.Config, jsonify jsonify.JSONify, logger logger.Logger,
	dbClient *gorm.DB, cache cache.Cache, validator validator.Validator) *UserMovieHandlerManager {
	// init movie usecase
	movieRepo := movieRepository.NewRepository(dbClient)
	movieCacheRepo := movieRepository.NewCacheRepository(cache)
	movieKinopoiskWebAPIRepo := movieRepository.NewKinopoiskWebAPIRepository(cfg, jsonify)
	movieUC := movieUsecase.NewUsecase(cfg, logger, movieRepo, movieCacheRepo, movieKinopoiskWebAPIRepo)
	// init user movie usecase
	userMovieRepo := repository.NewRepository(dbClient)
	userMovieUC := usecase.NewUsecase(userMovieRepo, movieUC)

	return &UserMovieHandlerManager{
		validator:   validator,
		userMovieUC: userMovieUC,
	}
}

// @summary		Получение информации о фильме
// @description	Получение информации о фильме по его kinopoisk ID
// @router			/films/full-info/{kinopoiskID} [get]
// @id				kinopoisk-get-film-info
// @tags			user-movie
// @security		JWT
// @param			kinopoiskID	path		int	true	"kinopoisk ID фильма"
// @success		200			{object}	entity.UserMovie
// @failure		401			"Пустой или неправильный токен"
// @failure		402			"Превышен дневной лимит запросов к Kinopoisk API"
// @failure		403			"Истекший или невалидный токен"
// @failure		404			"Фильм не найден"
// @failure		429			"Слишком много запросов. Лимит 5 запросов в секунду"
func (umhm UserMovieHandlerManager) GetUserMovie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(getFilmInfoIn)
		userMovie := new(entity.UserMovie)

		// parse path-params
		if err = ctx.ParamsParser(dataIn); err != nil {
			return fmt.Errorf("get movie info: %w", err)
		}
		// validate parsed data
		if err = umhm.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("get movie info: %w", err)
		}

		// add necessary data to userMovie
		// parse user ID from token from ctx
		userMovie.UserID, err = utils.ParseUserIDFromContext(ctx)
		if err != nil {
			return fmt.Errorf("get movie info: %w", err)
		}
		// init movie struct for user movie
		userMovie.Movie = &entity.Movie{KinopoiskID: dataIn.KinopoiskID}

		// find user movie with movie data (from DB or from API)
		err = umhm.userMovieUC.GetUserMovieByKinopoiskID(userMovie)
		if err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(userMovie)
	}
}

// @summary		Получение избранных фильмов юзера
// @description	Получение избранных фильмов юзера с пагинацией и настраиваемой сортировкой и фильтрацией
// @router			/films/stared [get]
// @id				films-get-stared
// @tags			user-movie
// @security		JWT
// @param			page		query		int			false	"страница поиска (Например: 1)"
// @param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
// @param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
// @param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
// @param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
// @param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
// @param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
// @param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
// @success		200			{object}	entity.UserMoviesWithCategory
// @failure		401			"Пустой или неправильный токен"
// @failure		403			"Истекший или невалидный токен"
func (umhm UserMovieHandlerManager) Stared() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userMoviesWithCategory := &entity.UserMoviesWithCategory{Category: getStaredConst}
		return umhm.getMoviesWithCategory(ctx, userMoviesWithCategory)
	}
}

// @summary		Получение фильмов юзера из списка "хочу посмотреть"
// @description	Получение фильмов юзера из списка "хочу посмотреть" с пагинацией и настраиваемой сортировкой и фильтрацией
// @router			/films/want [get]
// @id				films-get-want
// @tags			user-movie
// @security		JWT
// @param			page		query		int			false	"страница поиска (Например: 1)"
// @param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
// @param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
// @param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
// @param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
// @param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
// @param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
// @param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
// @success		200			{object}	entity.UserMoviesWithCategory
// @failure		401			"Пустой или неправильный токен"
// @failure		403			"Истекший или невалидный токен"
func (umhm UserMovieHandlerManager) Want() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userMoviesWithCategory := &entity.UserMoviesWithCategory{Category: getWantConst}
		return umhm.getMoviesWithCategory(ctx, userMoviesWithCategory)
	}
}

// @summary		Получение фильмов юзера из списка "посмотрел"
// @description	Получение фильмов юзера из списка "посмотрел" с пагинацией и настраиваемой сортировкой и фильтрацией
// @router			/films/watched [get]
// @id				films-get-watched
// @tags			user-movie
// @security		JWT
// @param			page		query		int			false	"страница поиска (Например: 1)"
// @param			sortField	query		string		false	"поле для сортировки (Например: year | По умолчанию: updated_at | Допустимые значения: title, rating, year, updated_at)"
// @param			sortOrder	query		string		false	"направление сортировки (Например: desc | По умолчанию: asc [для updated_at: desc] | Допустимые значения: asc, desc)"
// @param			ratingFrom	query		float64		false	"минимальный рейтинг (Например: 7.5 | Допустимые значения: 0 и больше)"
// @param			yearFrom	query		int			false	"минимальный год (Например: 1991 | Допустимые значения: 1500..3000)"
// @param			yearTo		query		int			false	"максимальный год (Например: 2014 | Допустимые значения: 1500..3000)"
// @param			type		query		string		false	"тип фильма (Например: сериал | Допустимые значения: фильм, сериал, видео, мини-сериал)"
// @param			genres		query		[]string	false	"жанры фильмов (Например: боевик)"
// @success		200			{object}	entity.UserMoviesWithCategory
// @failure		401			"Пустой или неправильный токен"
// @failure		403			"Истекший или невалидный токен"
func (umhm UserMovieHandlerManager) Watched() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userMoviesWithCategory := &entity.UserMoviesWithCategory{Category: getWatchedConst}
		return umhm.getMoviesWithCategory(ctx, userMoviesWithCategory)
	}
}

// get user movies from certain category
func (umhm UserMovieHandlerManager) getMoviesWithCategory(ctx *fiber.Ctx, userMoviesWithCategory *entity.UserMoviesWithCategory) error {
	var err error
	dataIn := new(categoryFilmsIn)

	// parse query-params
	if err = ctx.QueryParser(dataIn); err != nil {
		return fmt.Errorf("get movies from category %s: %w", userMoviesWithCategory.Category, err)
	}
	// validate parsed data
	if err = umhm.validator.Validate(dataIn); err != nil {
		return fmt.Errorf("get movies from category %s: %w", userMoviesWithCategory.Category, err)
	}

	// parse user ID from token from ctx
	userMoviesWithCategory.UserID, err = utils.ParseUserIDFromContext(ctx)
	if err != nil {
		return fmt.Errorf("get movies from category %s: %w", userMoviesWithCategory.Category, err)
	}
	// fill userMoviesWithCategory struct with parsed filter, sort and pagination settings
	userMoviesWithCategory.Filter = &dataIn.UserMoviesFilter
	userMoviesWithCategory.Sort = &dataIn.UserMoviesSort
	userMoviesWithCategory.Pagination = &dataIn.UserMoviesPagination

	// get user movies with given category and settings
	err = umhm.userMovieUC.GetUserMoviesWithCategory(userMoviesWithCategory)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(userMoviesWithCategory)
}
