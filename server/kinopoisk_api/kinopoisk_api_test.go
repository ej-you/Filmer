package kinopoisk_api

import (
	"fmt"
	"time"

	"testing"

	"server/db/schemas"
)


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)


var startTime time.Time

func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}

func successLog(t *testing.T, format string, a ...any) {
	t.Logf("\t%s\t%s", successMarker, fmt.Sprintf(format, a...))
}
func errorLog(t *testing.T, err error) {
	t.Logf("\t%s\tFailed: %v", failedMarker, err)
}


// search_films_by_keyword.go
func TestSearchFilmsByKeyword(t *testing.T) {
	startTime = time.Now()
	
	t.Log("Search films with keyword 'мстители'")
	{
		var films SearchedFilms

		err := SearchFilmsByKeyword("мстители", 1, &films)
		
		if err != nil {
			errorLog(t, err)
		} else {
			successLog(t, "Result films: %+v", films)
		}
	}
	logExecTime(t, &startTime)
}


// get_film_info.go
func TestGetFilmInfo(t *testing.T) {
	startTime = time.Now()
	
	t.Log("Get info about film 'Deadpool' with ID 462360")
	{
		var filmInfo schemas.RawFilmInfo

		err := GetFilmInfo(462360, &filmInfo)
		
		if err != nil {
			errorLog(t, err)
		} else {
			successLog(t, "Gotten film info: %+v", filmInfo)
		}
	}
	logExecTime(t, &startTime)
}


// get_film_staff.go
func TestGetFilmStaff(t *testing.T) {
	startTime = time.Now()
	
	t.Log("Get staff of film 'Deadpool' with ID 462360")
	{
		var filmStaff schemas.FilmStaff

		err := GetFilmStaff(462360, &filmStaff)
		
		if err != nil {
			errorLog(t, err)
		} else {
			successLog(t, "Gotten film staff: %+v", filmStaff)
		}
	}
	logExecTime(t, &startTime)
}
