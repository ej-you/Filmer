package http

import (
	"fmt"

	"gorm.io/gorm"
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/logger"
	"Filmer/server/pkg/validator"
	"Filmer/server/config"
	
	"Filmer/server/internal/movie"
	"Filmer/server/internal/movie/usecase"
	"Filmer/server/internal/movie/repository"
)


// Movie handlers manager
type MovieHandlerManager struct {
	validator	validator.Validator
    movieUC		movie.Usecase
}

// MovieHandlerManager constructor
func NewMovieHandlerManager(cfg *config.Config, jsonify jsonify.JSONify, logger logger.Logger, dbClient *gorm.DB,
	validator validator.Validator) *MovieHandlerManager {

	movieRepo := repository.NewRepository(dbClient)
	movieKinopoiskWebAPIRepo := repository.NewKinopoiskWebAPIRepository(cfg, jsonify)
	movieUC := usecase.NewUsecase(cfg, logger, movieRepo, movieKinopoiskWebAPIRepo)

	return &MovieHandlerManager{
		validator: validator,
		movieUC: movieUC,
	}
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
func (this MovieHandlerManager) SearchFilms() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		dataIn := new(searchFilmsIn)
		searchedMovies := new(entity.SearchedMovies)

		// parse query-params
		if err = ctx.QueryParser(dataIn); err != nil {
			return fmt.Errorf("search films: %w", err)
		}
		// validate parsed data
		if err = this.validator.Validate(dataIn); err != nil {
			return fmt.Errorf("search films: %w", err)
		}
		// get data from API
		searchedMovies.Query = dataIn.Query
		searchedMovies.Page = dataIn.Page
		if err = this.movieUC.SearchMovies(searchedMovies); err != nil {
			return err
		}
		return ctx.Status(200).JSON(searchedMovies)
	}
}
