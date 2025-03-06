package kinopoisk_api

import (
	"fmt"

	"server/db/schemas"
)


// получение полной информации о фильме по его ID
func GetFullFilmInfo(kinopoiskID int, outStruct *schemas.Movie) error {
	var err error

	// обращаемся к API за общей информацией о фильме
	if err = GetFilmInfo(kinopoiskID, outStruct); err != nil {
		return fmt.Errorf("get only film info: get full film info: %w", err)
	}
	// обращаемся к API за информацией о персонале фильма
	var filmStaff schemas.FilmStaff
	if err = GetFilmStaff(kinopoiskID, &filmStaff); err != nil {
		return fmt.Errorf("get only film staff info: get full film info: %w", err)
	}
	outStruct.Personal = &filmStaff
	return nil
}
