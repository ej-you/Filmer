// Package repository contains personal repositories implementations.
package repository

import (
	"fmt"
	"sort"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/staff"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/kinopoisk"
)

const (
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

var _ staff.KinopoiskRepo = (*kinopoiskRepo)(nil)

// staff.KinopoiskRepo implementation.
type kinopoiskRepo struct {
	cfg     *config.Config
	jsonify jsonify.JSONify
}

// Returns staff.KinopoiskRepo interface.
func NewKinopoiskRepo(cfg *config.Config, jsonify jsonify.JSONify) staff.KinopoiskRepo {
	return &kinopoiskRepo{
		cfg:     cfg,
		jsonify: jsonify,
	}
}

// Get full person info.
// Person ID must be presented.
// Fill given struct.
func (r kinopoiskRepo) GetFullInfoByID(person *entity.PersonFull) error {
	apiClient := kinopoisk.NewAPI(
		fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v1/staff/%d", person.ID),
		r.cfg.KinopoiskAPI.UnofficialKey,
		nil,
		r.jsonify,
	)
	// send request and process response
	rawPerson := &rawPersonFull{}
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
	person.Facts = rawPerson.Facts

	// filter all person movies to two slices by actor/director movie key
	allMoviesDirector, allMoviesActor := filterRawMovieList(rawPerson.Movies)
	// process filtered slices
	person.MoviesDirector = processFilteredMovieList(allMoviesDirector)
	person.MoviesActor = processFilteredMovieList(allMoviesActor)
	return nil
}

// Filter all person movies to two slices by director/actor movie key.
// Skip duplicates for each of the slice.
func filterRawMovieList(rawMovieList []rawPersonFullMovie) (
	allMoviesDirector []rawPersonFullMovie, allMoviesActor []rawPersonFullMovie) {

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
			// if movie with such ID not in director slice yet
			if _, exists = directorDupl[rawMovie.ID]; !exists {
				allMoviesDirector = append(allMoviesDirector, rawMovie)
				directorDupl[rawMovie.ID] = struct{}{}
			}
		case actorMoviesKey:
			// if movie with such ID not in actor slice yet
			if _, exists = actorDupl[rawMovie.ID]; !exists {
				allMoviesActor = append(allMoviesActor, rawMovie)
				actorDupl[rawMovie.ID] = struct{}{}
			}
		}
	}
	return allMoviesDirector, allMoviesActor
}

// Full processing of "actor" or "director" person movie list.
func processFilteredMovieList(filteredMovieList []rawPersonFullMovie) []entity.PersonFullMovie {
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
