package http

import (
	"fmt"
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/app/movie/repository"
	"Filmer/server/internal/app/movie/usecase"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/validator"
)

type MovieHandlerManager struct {
	validator validator.Validator
	movieUC   movie.Usecase
}

func NewMovieHandlerManager(cfg *config.Config, jsonify jsonify.JSONify, logger logger.Logger, dbClient *gorm.DB, cache cache.Storage,
	validator validator.Validator) *MovieHandlerManager {

	movieRepo := repository.NewDBRepo(dbClient)
	movieCacheRepo := repository.NewCacheRepo(cache, jsonify)
	movieKinopoiskWebAPIRepo := repository.NewKinopoiskRepo(cfg, jsonify)
	movieUC := usecase.NewUsecase(cfg, logger, movieRepo, movieCacheRepo, movieKinopoiskWebAPIRepo)

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
