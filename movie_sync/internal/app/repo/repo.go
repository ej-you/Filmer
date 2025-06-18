// Package repo contains interfaces of repositories for all entities.
// and its implementations like REST API.
package repo

import "Filmer/movie_sync/internal/app/entity"

type MovieAPIRepo interface {
	FullUpdate(movie *entity.Movie) error
}
