// Package repository contains personal repositories implementations.
package repository

import (
	"fmt"
	"sort"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	"Filmer/server/internal/personal"
	"Filmer/server/pkg/jsonify"
	kinopoiskAPI "Filmer/server/third_party/kinopoisk_api"
)

const (
	// searchMoviesPerPage = "30" // movies per page (for search movies)
	// movieDirectorsLimit = 8    // parse max 8 directors for movie

	// Person movies key for movies the director of which person was.
	directorMoviesKey = "DIRECTOR"
	// Person movies key for movies in which person was an actor.
	actorMoviesKey = "ACTOR"
	// Parse movies category for person if movies count is above than this value.
	personMoviesMinLimit = 10
	// Parse max 32 movies for person (for each of the "director" and "actor" categories).
	personMoviesMaxLimit = 32
)

// dict for converting person sex
var personSexMap = map[string]string{
	"MALE":   "мужской",
	"FEMALE": "женский",
}

var _ personal.KinopoiskWebAPIRepository = (*personalKinopoiskWebAPIRepository)(nil)

// personal.KinopoiskWebAPIRepository implementation.
type personalKinopoiskWebAPIRepository struct {
	cfg     *config.Config
	jsonify jsonify.JSONify
}

func NewKinopoiskWebAPIRepository(cfg *config.Config, jsonify jsonify.JSONify) personal.KinopoiskWebAPIRepository {
	return &personalKinopoiskWebAPIRepository{
		cfg:     cfg,
		jsonify: jsonify,
	}
}

// Get full person info.
// Person ID must be presented.
// Fill given struct.
func (p personalKinopoiskWebAPIRepository) GetFullInfoByID(person *entity.PersonFull) error {
	apiClient := kinopoiskAPI.NewKinopoiskAPI(
		fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v1/staff/%d", person.ID),
		p.cfg.KinopoiskAPI.UnofficialKey,
		nil,
		p.jsonify,
	)
	// send request and process response
	rawPerson := &entity.RawPersonFull{}
	err := apiClient.SendGET(rawPerson)
	if err != nil {
		return fmt.Errorf("get person info with kinopoisk API: %w", err)
	}

	// process received data
	person.Name = rawPerson.Name
	person.ImgURL = rawPerson.ImgURL
	person.Sex = personSexMap[rawPerson.Sex]
	person.Profession = rawPerson.Profession
	person.Age = rawPerson.Age
	person.Birthday = rawPerson.Birthday
	person.Death = rawPerson.Death

	// filter all person movies to two slices by actor/director movie key
	// var allMoviesDirector, allMoviesActor []entity.RawPersonFullMovie
	// for _, rawMovie := range rawPerson.Movies {
	// 	// skip unnamed movies
	// 	if rawMovie.Name == "" {
	// 		continue
	// 	}
	// 	switch rawMovie.ProfessionKey {
	// 	case directorMoviesKey:
	// 		allMoviesDirector = append(allMoviesDirector, rawMovie)
	// 	case actorMoviesKey:
	// 		allMoviesActor = append(allMoviesActor, rawMovie)
	// 	}
	// }
	allMoviesDirector, allMoviesActor := filterRawMovieList(rawPerson.Movies)

	person.MoviesDirector = processFilteredMovieList(allMoviesDirector)
	person.MoviesActor = processFilteredMovieList(allMoviesActor)
	return nil
}

// Filter all person movies to two slices by director/actor movie key.
// Skip duplicates for each of the slice.
func filterRawMovieList(rawMovieList []entity.RawPersonFullMovie) (
	allMoviesDirector []entity.RawPersonFullMovie, allMoviesActor []entity.RawPersonFullMovie) {

	directorDupl := make(map[int]struct{})
	actorDupl := make(map[int]struct{})

	var exists bool
	for _, rawMovie := range rawMovieList {
		// skip unnamed movies
		if rawMovie.Name == "" {
			continue
		}
		switch rawMovie.ProfessionKey {
		case directorMoviesKey:
			if _, exists = directorDupl[rawMovie.ID]; !exists {
				allMoviesDirector = append(allMoviesDirector, rawMovie)
				directorDupl[rawMovie.ID] = struct{}{}
			}
		case actorMoviesKey:
			if _, exists = actorDupl[rawMovie.ID]; !exists {
				allMoviesActor = append(allMoviesActor, rawMovie)
				actorDupl[rawMovie.ID] = struct{}{}
			}
		}
	}
	return allMoviesDirector, allMoviesActor
}

// Full processing of "actor" or "director" person movie list.
func processFilteredMovieList(filteredMovieList []entity.RawPersonFullMovie) []entity.PersonFullMovie {
	// skip movie list if movies count is too small
	if len(filteredMovieList) < personMoviesMinLimit {
		return nil
	}
	// sort movies slice by rating descending
	sort.Slice(filteredMovieList, func(i, j int) bool {
		return filteredMovieList[i].Rating > filteredMovieList[j].Rating
	})
	// cut to max size
	if len(filteredMovieList) > personMoviesMaxLimit {
		filteredMovieList = filteredMovieList[:personMoviesMaxLimit]
	}
	// process movies in list
	processedMovieList := make([]entity.PersonFullMovie, len(filteredMovieList))
	for i, rawMovie := range filteredMovieList {
		processedMovieList[i] = entity.PersonFullMovie{
			ID:    rawMovie.ID,
			Title: rawMovie.Name,
			Role:  rawMovie.Description,
		}
	}
	return processedMovieList
}

// // Get main movie info (without movie staff)
// // Fill given movie struct (apart of movie.Staff)
// func (mkr movieKinopoiskWebAPIRepository) getMovieInfoByKinopoiskID(movie *entity.Movie) error {
// 	apiClient := kinopoiskAPI.NewKinopoiskAPI(
// 		fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", movie.KinopoiskID),
// 		mkr.cfg.KinopoiskAPI.UnofficialKey,
// 		nil,
// 		mkr.jsonify,
// 	)
// 	// send request and process response
// 	rawFilmInfo := new(entity.RawMovieInfo)
// 	err := apiClient.SendGET(rawFilmInfo)
// 	if err != nil {
// 		return fmt.Errorf("get movie info with kinopoisk API: %w", err)
// 	}

// 	// process received data
// 	movie.KinopoiskID = rawFilmInfo.KinopoiskID
// 	movie.Title = rawFilmInfo.Title
// 	movie.ImgURL = rawFilmInfo.PosterURL
// 	movie.Rating = rawFilmInfo.RatingKinopoisk
// 	movie.WebURL = rawFilmInfo.WebURL
// 	movie.Year = rawFilmInfo.Year
// 	movie.Description = rawFilmInfo.Description
// 	movie.Genres = rawFilmInfo.Genres
// 	// if movie length was not found
// 	if rawFilmInfo.FilmLenMinutes == 0 {
// 		movie.MovieLength = ""
// 	} else {
// 		movie.MovieLength = utils.RawMinutesToTime(rawFilmInfo.FilmLenMinutes)
// 	}
// 	// "фильм" by default
// 	movie.Type = movieTypesMap[rawFilmInfo.Type]
// 	if movie.Type == "" {
// 		movie.Type = "фильм"
// 	}
// 	return nil
// }

// // Get info about the movie staff
// // Fill given movieStaff struct
// func (mkr movieKinopoiskWebAPIRepository) getMovieStaffByMovieKinopoiskID(movieKinopoiskID int, movieStaff *entity.MovieStaff) error {
// 	apiClient := kinopoiskAPI.NewKinopoiskAPI(
// 		"https://kinopoiskapiunofficial.tech/api/v1/staff",
// 		mkr.cfg.KinopoiskAPI.UnofficialKey,
// 		map[string]string{
// 			"filmId": fmt.Sprint(movieKinopoiskID),
// 		},
// 		mkr.jsonify,
// 	)

// 	var rawFilmStaffSlice entity.RawMovieStaffSlice
// 	// send request and process response
// 	err := apiClient.SendGET(&rawFilmStaffSlice)
// 	if err != nil {
// 		return fmt.Errorf("get movie staff with kinopoisk API: %w", err)
// 	}

// 	// init slices for movie staff
// 	movieStaff.Directors = make([]entity.Person, 0)
// 	movieStaff.Actors = make([]entity.Person, 0, movieActorsLimit)
// 	// sort movie staff to slices
// 	for _, rawFilmStaff := range rawFilmStaffSlice {
// 		switch rawFilmStaff.ProfessionKey {
// 		case "DIRECTOR":
// 			// max movieDirectorsLimit actors for movie
// 			if len(movieStaff.Directors) == movieDirectorsLimit {
// 				continue
// 			}
// 			movieStaff.Directors = append(movieStaff.Directors, entity.Person{
// 				ID:     rawFilmStaff.StaffID,
// 				Name:   rawFilmStaff.Name,
// 				ImgURL: rawFilmStaff.ImgURL,
// 			})
// 		case "ACTOR":
// 			// max movieActorsLimit actors for movie
// 			if len(movieStaff.Actors) == movieActorsLimit {
// 				continue
// 			}
// 			movieStaff.Actors = append(movieStaff.Actors, entity.Person{
// 				ID:     rawFilmStaff.StaffID,
// 				Name:   rawFilmStaff.Name,
// 				Role:   &rawFilmStaff.Description,
// 				ImgURL: rawFilmStaff.ImgURL,
// 			})
// 		}
// 	}
// 	return nil
// }
