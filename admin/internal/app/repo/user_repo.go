package repo

import (
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"

	"Filmer/admin/internal/app/entity"
)

const (
	_requestTimeout = 2 * time.Second        // timeout for requests to REST API
	_retryCount     = 3                      // amount of retries attempts in error cases
	_retryInitTime  = 500 * time.Millisecond // time between first request and first retry
	_retryMaxTime   = 2 * time.Second        // max time between request and retry
)

var _ UserAPIRepo = (*userAPIRepo)(nil)

// UserAPIRepo implementation.
type userAPIRepo struct {
	host string
}

func NewUserAPIRepo(host string) UserAPIRepo {
	return &userAPIRepo{
		host: host,
	}
}

// GetUsersActivity sends request to REST API and gets JSON-response with slice of users activity.
func (r *userAPIRepo) GetUsersActivity() (entity.UsersActivity, error) {
	var result entity.UsersActivity
	// init client with retry params
	client := resty.New().
		SetTimeout(_requestTimeout).
		SetRetryCount(_retryCount).
		SetRetryWaitTime(_retryInitTime).
		SetRetryMaxWaitTime(_retryMaxTime)
	// do request to REST API and parse JSON-response to result
	_, err := client.R().
		SetResult(&result).
		Get(r.host + "/api/v1/user/all/activity")
	if err != nil {
		return nil, fmt.Errorf("resuest to rest api: %w", err)
	}
	return result, nil
}
