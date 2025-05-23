package repository

import (
	"fmt"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/movie"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/kinopoisk"
	"Filmer/server/internal/pkg/utils"
)

const (
	searchMoviesPerPage = "30" // movies per page (for search movies)
	movieDirectorsLimit = 8    // parse max 8 directors for movie
	movieActorsLimit    = 30   // parse max 30 actors for movie
)

// dict for converting movie types
var movieTypesMap = map[string]string{
	"FILM":        "фильм",
	"TV_SERIES":   "сериал",
	"VIDEO":       "видео",
	"MINI_SERIES": "мини-сериал",
	"TV_SHOW":     "сериал",
}

var _ movie.KinopoiskRepo = (*kinopoiskRepo)(nil)

// movie.KinopoiskRepo implementation.
type kinopoiskRepo struct {
	cfg     *config.Config
	jsonify jsonify.JSONify
}

// Returns movie.KinopoiskRepo interface.
func NewKinopoiskRepo(cfg *config.Config, jsonify jsonify.JSONify) movie.KinopoiskRepo {
	return &kinopoiskRepo{
		cfg:     cfg,
		jsonify: jsonify,
	}
}

// Search movies by keyword.
// Must be presented query (searchedMovies.Query) and page (searchedMovies.Page).
// Fill given searchedMovies struct.
func (r kinopoiskRepo) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	apiClient := kinopoisk.NewKinopoiskAPI(
		"https://api.kinopoisk.dev/v1.4/movie/search",
		r.cfg.KinopoiskAPI.Key,
		map[string]string{
			"query": searchedMovies.Query,
			"page":  fmt.Sprint(searchedMovies.Page),
			"limit": searchMoviesPerPage,
		},
		r.jsonify,
	)
	// send request and process response
	if err := apiClient.SendGET(searchedMovies); err != nil {
		return fmt.Errorf("search movies with kinopoisk API: %w", err)
	}
	return nil
}

// Get full movie info (including movie staff).
// Must be presented kinopoisk movie ID (movie.KinopoiskID).
// Fill given movie struct.
func (r kinopoiskRepo) GetFullMovieByKinopoiskID(movie *entity.Movie) error {
	var err error

	if err = r.getMovieInfoByKinopoiskID(movie); err != nil {
		return fmt.Errorf("get full movie: %w", err)
	}
	// init staff struct for movie
	movie.Staff = new(entity.MovieStaff)
	if err = r.getMovieStaffByMovieKinopoiskID(movie.KinopoiskID, movie.Staff); err != nil {
		return fmt.Errorf("get full movie: %w", err)
	}
	return nil
}

// Get main movie info (without movie staff).
// Fill given movie struct (apart of movie.Staff).
func (r kinopoiskRepo) getMovieInfoByKinopoiskID(movie *entity.Movie) error {
	apiClient := kinopoisk.NewKinopoiskAPI(
		fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", movie.KinopoiskID),
		r.cfg.KinopoiskAPI.UnofficialKey,
		nil,
		r.jsonify,
	)
	// send request and process response
	rawFilmInfo := new(entity.RawMovieInfo)
	err := apiClient.SendGET(rawFilmInfo)
	if err != nil {
		return fmt.Errorf("get movie info with kinopoisk API: %w", err)
	}

	// process received data
	movie.KinopoiskID = rawFilmInfo.KinopoiskID
	movie.Title = rawFilmInfo.Title
	movie.ImgURL = rawFilmInfo.PosterURL
	movie.Rating = rawFilmInfo.RatingKinopoisk
	movie.WebURL = rawFilmInfo.WebURL
	movie.Year = rawFilmInfo.Year
	movie.Description = rawFilmInfo.Description
	movie.Genres = rawFilmInfo.Genres
	// if movie length was not found
	if rawFilmInfo.FilmLenMinutes == 0 {
		movie.MovieLength = ""
	} else {
		movie.MovieLength = utils.RawMinutesToTime(rawFilmInfo.FilmLenMinutes)
	}
	// "фильм" by default
	movie.Type = movieTypesMap[rawFilmInfo.Type]
	if movie.Type == "" {
		movie.Type = "фильм"
	}
	return nil
}

// Get info about the movie staff.
// Fill given movieStaff struct.
func (r kinopoiskRepo) getMovieStaffByMovieKinopoiskID(movieKinopoiskID int, movieStaff *entity.MovieStaff) error {
	apiClient := kinopoisk.NewKinopoiskAPI(
		"https://kinopoiskapiunofficial.tech/api/v1/staff",
		r.cfg.KinopoiskAPI.UnofficialKey,
		map[string]string{
			"filmId": fmt.Sprint(movieKinopoiskID),
		},
		r.jsonify,
	)

	var rawFilmStaffSlice entity.RawMovieStaffSlice
	// send request and process response
	err := apiClient.SendGET(&rawFilmStaffSlice)
	if err != nil {
		return fmt.Errorf("get movie staff with kinopoisk API: %w", err)
	}

	// init slices for movie staff
	movieStaff.Directors = make([]entity.Person, 0)
	movieStaff.Actors = make([]entity.Person, 0, movieActorsLimit)
	// sort movie staff to slices
	for _, rawFilmStaff := range rawFilmStaffSlice {
		switch rawFilmStaff.ProfessionKey {
		case "DIRECTOR":
			// max movieDirectorsLimit actors for movie
			if len(movieStaff.Directors) == movieDirectorsLimit {
				continue
			}
			movieStaff.Directors = append(movieStaff.Directors, entity.Person{
				ID:     rawFilmStaff.StaffID,
				Name:   rawFilmStaff.Name,
				ImgURL: rawFilmStaff.ImgURL,
			})
		case "ACTOR":
			// max movieActorsLimit actors for movie
			if len(movieStaff.Actors) == movieActorsLimit {
				continue
			}
			movieStaff.Actors = append(movieStaff.Actors, entity.Person{
				ID:     rawFilmStaff.StaffID,
				Name:   rawFilmStaff.Name,
				Role:   &rawFilmStaff.Description,
				ImgURL: rawFilmStaff.ImgURL,
			})
		}
	}
	return nil
}
