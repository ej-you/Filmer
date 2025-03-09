package repository

import (
	"fmt"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/utils"
	kinopoiskAPI "Filmer/server/third_party/kinopoisk_api"
	"Filmer/server/config"
	
	"Filmer/server/internal/movie"
)


const searchMoviesPerPage = "25" // movies per page (for search movies)

// dict for converting movie types
var movieTypesMap = map[string]string {
	"FILM": "фильм",
	"TV_SERIES": "сериал",
	"VIDEO": "видео",
	"MINI_SERIES": "мини-сериал",
	"TV_SHOW": "сериал",
}


// movie.KinopoiskWebAPIRepository interface implementation
type movieKinopoiskWebAPIRepository struct {
	cfg 	*config.Config
	jsonify jsonify.JSONify
}

// movie.KinopoiskWebAPIRepository constructor
func NewKinopoiskWebAPIRepository(cfg *config.Config, jsonify jsonify.JSONify) movie.KinopoiskWebAPIRepository {
	return &movieKinopoiskWebAPIRepository{
		cfg: cfg,
		jsonify: jsonify,
	}
}

// Search movies by keyword
// Fill given searchedMovies struct
func (this movieKinopoiskWebAPIRepository) SearchMovies(searchedMovies *entity.SearchedMovies) error {
	apiClient := kinopoiskAPI.NewKinopoiskAPI(
		"https://api.kinopoisk.dev/v1.4/movie/search",
		this.cfg.KinopoiskAPI.Key,
		map[string]string{
			"query": searchedMovies.Query,
			"page": fmt.Sprint(searchedMovies.Page),
			"limit": searchMoviesPerPage,
		},
		this.jsonify,
	)
	// send request and process response
	if err := apiClient.SendGET(searchedMovies); err != nil {
		return fmt.Errorf("search movies with kinopoisk API: %w", err)
	}
	return nil
}

// Get full movie info (including movie staff)
// Fill given movie struct
func (this movieKinopoiskWebAPIRepository) GetFullMovieByKinopoiskID(movie *entity.Movie) error {
	var err error

	if err = this.getMovieInfoByKinopoiskID(movie); err != nil {
		return fmt.Errorf("get full movie: %w", err)
	}
	if err = this.getMovieStaffByMovieKinopoiskID(movie.KinopoiskID, movie.Staff); err != nil {
		return fmt.Errorf("get full movie: %w", err)
	}
	return nil
}

// Get main movie info (without movie staff)
// Fill given movie struct (apart of movie.Staff)
func (this movieKinopoiskWebAPIRepository) getMovieInfoByKinopoiskID(movie *entity.Movie) error {
	apiClient := kinopoiskAPI.NewKinopoiskAPI(
		fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", movie.KinopoiskID),
		this.cfg.KinopoiskAPI.UnofficialKey,
		nil,
		this.jsonify,
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

// Get info about the movie staff
// Fill given movieStaff struct
func (this movieKinopoiskWebAPIRepository) getMovieStaffByMovieKinopoiskID(movieKinopoiskID int, movieStaff *entity.MovieStaff) error {
	apiClient := kinopoiskAPI.NewKinopoiskAPI(
		"https://kinopoiskapiunofficial.tech/api/v1/staff",
		this.cfg.KinopoiskAPI.UnofficialKey,
		map[string]string{
			"filmId": fmt.Sprint(movieKinopoiskID),
		},
		this.jsonify,
	)

	var rawFilmStaffSlice entity.RawMovieStaffSlice
	// send request and process response
	err := apiClient.SendGET(&rawFilmStaffSlice)
	if err != nil {
		return fmt.Errorf("get movie staff with kinopoisk API: %w", err)
	}

	// init slices for movie staff
	movieStaff.Directors = make([]entity.Person, 0)
	movieStaff.Actors = make([]entity.Person, 0, 30)
	// sort movie staff to slices
	for _, rawFilmStaff := range rawFilmStaffSlice {
		switch rawFilmStaff.ProfessionKey {
			case "DIRECTOR":
				movieStaff.Directors = append(movieStaff.Directors, entity.Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
			case "ACTOR":
				// max 30 actors for movie
				if len(movieStaff.Actors) == 30 {
					continue
				}
				movieStaff.Actors = append(movieStaff.Actors, entity.Person{
					ID: rawFilmStaff.StaffID,
					Name: rawFilmStaff.Name,
					Role: &rawFilmStaff.Description,
					ImgUrl: rawFilmStaff.ImgUrl,
				})
		}
	}
	return nil
}
