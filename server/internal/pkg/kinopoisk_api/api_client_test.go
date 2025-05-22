package kinopoisk_api

import (
	"testing"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/pkg/jsonify"
)

func TestKinopoiskAPI(t *testing.T) {
	t.Log("Try to get film staff from Kinopoisk API")

	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	jsonify := jsonify.NewJSONify()

	url := "https://kinopoiskapiunofficial.tech/api/v1/staff"
	queryParams := map[string]string{"filmId": "301"}
	kinopoiskAPI := NewKinopoiskAPI(url, cfg.KinopoiskAPI.UnofficialKey, queryParams, jsonify)

	var rawFilmStaff entity.RawMovieStaffSlice

	if err := kinopoiskAPI.SendGET(&rawFilmStaff); err != nil {
		t.Error("ERROR:", err)
		return
	}
	t.Log("rawFilmStaff:", rawFilmStaff)
}
