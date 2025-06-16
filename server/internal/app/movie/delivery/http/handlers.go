package http

import (
	"fmt"
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	kinopoiskrepo "Filmer/server/internal/app/kinopoisk/repository"
	kinopoiskusecase "Filmer/server/internal/app/kinopoisk/usecase"
	"Filmer/server/internal/app/movie"
	movierepo "Filmer/server/internal/app/movie/repository"
	movieusecase "Filmer/server/internal/app/movie/usecase"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/validator"
)

type MovieHandlerManager struct {
	validator validator.Validator
	movieUC   movie.Usecase
}

func NewMovieHandlerManager(cfg *config.Config, jsonify jsonify.JSONify, logger logger.Logger,
	dbClient *gorm.DB, cache cache.Storage, validator validator.Validator) *MovieHandlerManager {

	// init kinopoisk usecase
	kinopoiskCacheRepo := kinopoiskrepo.NewCacheRepo(cache)
	kinopoiskUC := kinopoiskusecase.NewUsecase(kinopoiskCacheRepo)
	// init movie usecase
	movieRepo := movierepo.NewDBRepo(dbClient)
	movieKinopoiskWebAPIRepo := movierepo.NewKinopoiskRepo(cfg, jsonify)
	movieUC := movieusecase.NewUsecase(cfg, logger, movieRepo,
		movieKinopoiskWebAPIRepo, kinopoiskUC)

	return &MovieHandlerManager{
		validator: validator,
		movieUC:   movieUC,
	}
}

// @summary		Поиск фильмов
// @description	Поиск фильмов по ключевому слову с пагинацией
// @router			/kinopoisk/films/search [get]
// @id				kinopoisk-search-films
// @tags			movie
// @security		JWT
// @param			q		query		string	true	"ключевое слово (Например: матрица)"
// @param			page	query		int		true	"страница поиска (Например: 1)"
// @success		200		{object}	entity.SearchedMovies
// @failure		401		"Пустой или неправильный токен"
// @failure		402		"Превышен дневной лимит запросов к Kinopoisk API"
// @failure		403		"Истекший или невалидный токен"
// @failure		404		"Фильмы не найдены"
// @failure		429		"Слишком много запросов. Лимит 5 запросов в секунду"
func (h MovieHandlerManager) SearchFilms() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(searchFilmsIn)
		searchedMovies := new(entity.SearchedMovies)

		// parse query-params
		if err = ctx.QueryParser(dataIn); err != nil {
			return fmt.Errorf("search films: %w", err)
		}
		// validate parsed data
		if err = h.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("search films: %w", err)
		}
		// get data from API
		searchedMovies.Query = strings.ToLower(dataIn.Query)
		searchedMovies.Page = dataIn.Page
		err = h.movieUC.SearchMovies(searchedMovies)
		if err != nil {
			return err
		}
		return ctx.Status(http.StatusOK).JSON(searchedMovies)
	}
}

// @summary		Обновление информации о фильме
// @description	Полное обновление информации о фильме из Kinopoisk API.
// @router			/kinopoisk/films/update-movie/{movieID} [post]
// @id				kinopoisk-update-movie
// @tags			movie
// @param			movieID	path	string	true	"UUID фильма из БД"
// @success		204		"No Content"
// @failure		402		"Превышен дневной лимит запросов к Kinopoisk API"
// @failure		404		"Фильм не найден"
func (h MovieHandlerManager) UpdateMovie() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dataIn := &updateMovieIn{}
		movie := &entity.Movie{}

		// parse path-params
		if err := ctx.ParamsParser(dataIn); err != nil {
			return fmt.Errorf("update movie: %w", err)
		}
		// validate parsed data
		if err := h.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("update movie: %w", err)
		}
		movie.ID = dataIn.MovieID
		// update movie info
		if err := h.movieUC.FullUpdate(movie); err != nil {
			return err
		}
		return ctx.Status(http.StatusNoContent).Send(nil)
	}
}
