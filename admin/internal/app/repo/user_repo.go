package repo

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"

	"Filmer/admin/internal/app/entity"
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
	// init client
	client := resty.New()
	// do request to REST API and parse JSON-response to result
	_, err := client.R().
		SetResult(&result).
		Get(r.host + "/api/v1/user/all/activity")
	if err != nil {
		return nil, fmt.Errorf("resuest to rest api: %w", err)
	}
	return result, nil
}
