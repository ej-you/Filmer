// Package personal contains usecase and repositories interfaces for personal entities.
// Subpackage usecase contains usecase implementation.
// Subpackage repository contains repositories implementations.
// Subpackage delivery/http contains http router and handlers for personal usecase.
package personal

import (
	"Filmer/server/internal/app/entity"
)

type Usecase interface {
	GetByID(person *entity.PersonFull) error
}
