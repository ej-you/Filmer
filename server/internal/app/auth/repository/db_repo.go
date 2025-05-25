package repository

import (
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/pkg/httperror"

	"Filmer/server/internal/app/auth"
)

var _ auth.DBRepo = (*dbRepo)(nil)

// auth.DBRepo implementation.
type dbRepo struct {
	dbClient *gorm.DB
}

// Returns auth.DBRepo interface.
func NewDBRepo(dbClient *gorm.DB) auth.DBRepo {
	return &dbRepo{
		dbClient: dbClient,
	}
}

// Create new user.
// Fill given user struct.
func (r dbRepo) CreateUser(user *entity.User) error {
	createResult := r.dbClient.Create(user)
	if err := createResult.Error; err != nil {
		// if user with such email already exists
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return httperror.New(http.StatusConflict, "user with such email already exists", err)
		}
		return httperror.New(http.StatusInternalServerError, "failed to create user", err)
	}
	return nil
}

// Get user by email.
// User email (user.Email) must be presented.
// Fill given user struct.
// Returns error even if user not found.
func (r dbRepo) GetUserByEmail(user *entity.User) error {
	selectResult := r.dbClient.Where("email = ?", user.Email).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httperror.New(http.StatusNotFound, "user with such email was not found", err)
		}
		return httperror.New(http.StatusInternalServerError, "failed to get user by email", err)
	}
	return nil
}
