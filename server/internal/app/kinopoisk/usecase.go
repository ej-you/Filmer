// Package kinopoisk contains main usecase and repositories interfaces for all kinopoisk API.
// Subpackage usecase contains usecase implementation.
// Subpackage repository contains repositories implementations.
package kinopoisk

type Usecase interface {
	SetOfficialAPILimit() error
	IsOfficialAPILimitReached() (bool, error)

	SetUnofficialAPILimit() error
	IsUnofficialAPILimitReached() (bool, error)
}
