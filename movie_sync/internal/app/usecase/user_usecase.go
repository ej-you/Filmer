package usecase

import (
	"fmt"

	"Filmer/movie_sync/internal/app/entity"
	"Filmer/movie_sync/internal/app/repo"
)

var _ MovieUsecase = (*movieUsecase)(nil)

// MovieUsecase implementation.
type movieUsecase struct {
	movieAPIRepo repo.MovieAPIRepo
}

func NewMovieUsecase(movieAPIRepo repo.MovieAPIRepo) MovieUsecase {
	return &movieUsecase{
		movieAPIRepo: movieAPIRepo,
	}
}

// FullUpdate updates movie info.
// Movie id must be set (via NewMovie constructor).
func (u *movieUsecase) FullUpdate(movie *entity.Movie) error {
	err := u.movieAPIRepo.FullUpdate(movie)
	if err != nil {
		return fmt.Errorf("full update movie: %w", err)
	}
	return nil
}
