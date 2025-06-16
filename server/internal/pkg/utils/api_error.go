package utils

import (
	"errors"
	"net/http"

	"Filmer/server/internal/app/kinopoisk"
	"Filmer/server/internal/pkg/httperror"
)

// HandleOfficialAPIError sets official api limit value to cache
// if given error is 402 payment required error.
// Returns cache error if it occurs, else given error.
func HandleOfficialAPIError(err error, kinopoiskUC kinopoisk.Usecase) error {
	var httpErr httperror.HTTPError
	// assert gotten error to http error AND check gotten http error is 402
	if errors.As(err, &httpErr) && httpErr.StatusCode() == http.StatusPaymentRequired {
		// set limit is reached to cache
		if err := kinopoiskUC.SetOfficialAPILimit(); err != nil {
			return err
		}
	}
	return err
}

// HandleUnofficialAPIError sets unofficial api limit value to cache
// if given error is 402 payment required error.
// Returns cache error if it occurs, else given error.
func HandleUnofficialAPIError(err error, kinopoiskUC kinopoisk.Usecase) error {
	var httpErr httperror.HTTPError
	// assert gotten error to http error AND check gotten http error is 402
	if errors.As(err, &httpErr) && httpErr.StatusCode() == http.StatusPaymentRequired {
		// set limit is reached to cache
		if err := kinopoiskUC.SetUnofficialAPILimit(); err != nil {
			return err
		}
	}
	return err
}
