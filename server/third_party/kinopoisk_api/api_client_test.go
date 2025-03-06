package kinopoisk_api

import (
	"testing"

	"Filmer/server/internal/entity"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/config"
)


func TestKinopoiskAPI(t *testing.T) {
	t.Log("Try to get film staff from Kinopoisk API")

	cfg := config.NewConfig()
	jsonify := jsonify.NewJSONify()

	url := "https://kinopoiskapiunofficial.tech/api/v1/staff"
	queryParams := map[string]string{"filmId": "301"}
	kinopoiskAPI := NewKinopoiskAPI(url, cfg.KinopoiskAPI.UnofficialKey, queryParams, jsonify)

	var rawFilmStaff entity.RawFilmStaffSlice

	err := kinopoiskAPI.SendGET(&rawFilmStaff)
	if err != nil {
		t.Error("ERROR:", err)
		return
	}
	t.Log("rawFilmStaff:", rawFilmStaff)
}
