package kinopoisk_api

import "fmt"


// перевод минут в часы:минуты
func RawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/60, minutes%60)
}

// перевод списка структур с жанрами с список жанров
func RawGenresToSlice(genres []GenreUnofficial) []string {
	genresSlice := make([]string, len(genres), len(genres))

	for i, genre := range genres {
		genresSlice[i] = genre.Genre
	}
	return genresSlice
}
