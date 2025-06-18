// Package usecase contains interfaces of usecases
// and its implementations for all entities.
package usecase

import "Filmer/movie_sync/internal/app/entity"

type MovieUsecase interface {
	FullUpdate(movie *entity.Movie) error
}
