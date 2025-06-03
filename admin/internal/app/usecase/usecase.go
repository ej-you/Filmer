// Package usecase contains interfaces of usecases
// and its implementations for all entities.
package usecase

import "Filmer/admin/internal/app/entity"

type UserUsecase interface {
	GetUsersActivity() (entity.UsersActivity, error)
}
