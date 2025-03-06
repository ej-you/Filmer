package background

import (
	"github.com/google/uuid"

	kinopoiskAPI "server/kinopoisk_api"
	"server/db/schemas"
	"server/db"
	"server/settings"
)


// фоновое обновление информации о фильме
func UpdateFilmData(movieID uuid.UUID, kinopoiskID int) {
	var err error
	var movie schemas.Movie

	// выставляем ID фильма (первичный ключ) для обновления записи, а не создания новой
	movie.ID = movieID
	// обращаемся к API
	if err = kinopoiskAPI.GetFullFilmInfo(kinopoiskID, &movie); err != nil {
		settings.ErrorLog.Printf("background update for film %d: failed to get film info from API: %v", kinopoiskID, err)
		return
	}
	// обновляем фильм в БД
	createResult := db.GetConn().Save(&movie)
	if err = createResult.Error; err != nil {
		settings.ErrorLog.Printf("background update for film %d: failed to save movie: %v", kinopoiskID, err)
		return
	}
	settings.InfoLog.Printf("Background update for film %d", kinopoiskID)
}
