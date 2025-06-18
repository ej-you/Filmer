package repo

import (
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"

	"Filmer/movie_sync/internal/app/entity"
)

const (
	_requestTimeout = 2 * time.Second        // timeout for requests to REST API
	_retryCount     = 3                      // amount of retries attempts in error cases
	_retryInitTime  = 500 * time.Millisecond // time between first request and first retry
	_retryMaxTime   = 2 * time.Second        // max time between request and retry
)

var _ MovieAPIRepo = (*movieAPIRepo)(nil)

// MovieAPIRepo implementation.
type movieAPIRepo struct {
	host string
}

func NewMovieAPIRepo(host string) MovieAPIRepo {
	return &movieAPIRepo{
		host: host,
	}
}

// FullUpdate sends request to REST API for full update movie info.
// Movie id must be set (via NewMovie constructor).
func (r *movieAPIRepo) FullUpdate(movie *entity.Movie) error {
	// init client with retry params
	client := resty.New().
		SetTimeout(_requestTimeout).
		SetRetryCount(_retryCount).
		SetRetryWaitTime(_retryInitTime).
		SetRetryMaxWaitTime(_retryMaxTime)
		// TODO: try SetError
	// do request to REST API
	_, err := client.R().
		Post(r.host + "/api/v1/kinopoisk/films/update-movie/" + movie.ID())
	if err != nil {
		return fmt.Errorf("resuest to rest api: %w", err)
	}
	return nil
}
