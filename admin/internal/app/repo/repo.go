// Package repo contains interfaces of repositories for all entities.
// and its implementations like REST API.
package repo

import "Filmer/admin/internal/app/entity"

type UserAPIRepo interface {
	GetUsersActivity() (entity.UsersActivity, error)
}
