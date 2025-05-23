// Package staff contains usecase and repositories interfaces for person/staff entities.
// Subpackage usecase contains usecase implementation.
// Subpackage repository contains repositories implementations.
// Subpackage delivery/http contains http router and handlers for personal/staff usecase.
package staff

import (
	"Filmer/server/internal/app/entity"
)

type Usecase interface {
	GetByID(person *entity.PersonFull) error
}
