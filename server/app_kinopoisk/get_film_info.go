package app_kinopoisk

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	fiber "github.com/gofiber/fiber/v2"

	kinopoiskAPI "server/kinopoisk_api"
	coreValidator "server/core/validator"
	"server/background"
	"server/core/services"
	"server/db/schemas"
	"server/db"
	"server/settings"
)


//easyjson:json
// данные для получения информации о фильме
type GetFilmInfoIn struct {
	// kinopoisk ID фильма
	KinopoiskID int `params:"kinopoiskID" validate:"required"`
}


//	@summary		Получение информации о фильме
//	@description	Получение информации о фильме по его kinopoisk ID
//	@router			/kinopoisk/films/{kinopoiskID} [get]
//	@id				kinopoisk-get-film-info
//	@tags			kinopoisk-films
//	@security		JWT
//	@param			kinopoiskID	path		int	true	"kinopoisk ID фильма"
//	@success		200			{object}	schemas.UserMovie
//	@failure		401			"Пустой или неправильный токен"
//	@failure		402			"Превышен дневной лимит запросов к Kinopoisk API"
//	@failure		403			"Истекший или невалидный токен"
//	@failure		404			"Фильм не найден"
//	@failure		429			"Слишком много запросов. Лимит 5 запросов в секунду"
func GetFilmInfo(ctx *fiber.Ctx) error {
	var err error
	var dataIn GetFilmInfoIn
	var userMovieInfo schemas.UserMovie

	// парсинг path-параметров
	if err = ctx.ParamsParser(&dataIn); err != nil {
		return fmt.Errorf("get film info: %w", err)
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Validate(&dataIn); err != nil {
		return fmt.Errorf("get film info: %w", err)
	}
	
	// поиск фильма в БД
	found, err := getFilmInfoFromDB(services.ParseUserIDFromContext(ctx), &dataIn, &userMovieInfo)
	if err != nil {
		return fmt.Errorf("film %d: get info: %w", dataIn.KinopoiskID, err)
	}

	// если фильм найден, то возвращаем его
	if found {
		return ctx.JSON(userMovieInfo)
	}
	// иначе обращаемся к API
	if err = kinopoiskAPI.GetFullFilmInfo(dataIn.KinopoiskID, userMovieInfo.Movie); err != nil {
		return err
	}
	// сохраняем фильм в БД
	createResult := db.GetConn().Create(&userMovieInfo.Movie)
	if err = createResult.Error; err != nil {
		return fmt.Errorf("film %d: get info: %w", dataIn.KinopoiskID, fiber.NewError(500, "failed to save movie: " + err.Error()))
	}

	settings.InfoLog.Printf("Got info about film %d from Kinopoisk API", dataIn.KinopoiskID)
	return ctx.Status(200).JSON(userMovieInfo)
}


// получение информации о фильме из БД (true, если запись найдена)
func getFilmInfoFromDB(userID uuid.UUID, dataIn *GetFilmInfoIn, userMovie *schemas.UserMovie) (bool, error) {
	var err error
	dbConn := db.GetConn()
	
	// поиск в БД
	selectResult := dbConn.
		Where("kinopoisk_id = ?", dataIn.KinopoiskID).
		Preload("Genres").
		First(&userMovie.Movie)
	if err = selectResult.Error; err != nil {
		// если ошибка НЕ в ненахождении записи
		if !errors.Is(err, db.NotFoundError) {
			return false, fmt.Errorf("film %d: get info: %w", dataIn.KinopoiskID, fiber.NewError(500, "failed to get movie: " + err.Error()))
		}
		// если запись не найдена
		return false, nil
	}

	expiredAt := userMovie.Movie.UpdatedAt.Add(settings.KinopoiskDataUpdateDuration).UTC()
	now := time.Now().UTC()
	// если данные из БД устарели, то запускаем их обновление в фоне
	if now.After(expiredAt) {
		go background.UpdateFilmData(userMovie.Movie.ID, userMovie.Movie.KinopoiskID)
	}

	// ищем фильм в таблице фильмов юзеров (для подгрузки статуса и метки избранного)
	selectResult = dbConn.
		Where("movie_id = ? AND user_id = ?", userMovie.Movie.ID, userID).
		First(userMovie)
	if err = selectResult.Error; err != nil {
		// если ошибка НЕ в ненахождении записи
		if !errors.Is(err, db.NotFoundError) {
			return false, fmt.Errorf("film %d: get info: %w", dataIn.KinopoiskID, fiber.NewError(500, "failed to get user movie: " + err.Error()))
		}
	}
	return true, nil
}
