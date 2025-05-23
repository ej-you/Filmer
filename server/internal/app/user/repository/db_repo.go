package repository

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/user"
	"Filmer/server/internal/pkg/httperror"
)

var _ user.DBRepo = (*dbRepo)(nil)

// user.DBRepo implementation.
type dbRepo struct {
	dbClient *gorm.DB
}

// Returns user.DBRepo interface.
func NewDBRepo(dbClient *gorm.DB) user.DBRepo {
	return &dbRepo{
		dbClient: dbClient,
	}
}

// Get user by ID.
// User ID (user.ID) must be presented.
// Fill given user struct.
// Returns error even if user not found.
func (r dbRepo) GetUserByID(user *entity.User) error {
	selectResult := r.dbClient.Where("id = ?", user.ID).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httperror.NewHTTPError(http.StatusNotFound, "user with such id was not found", err)
		}
		return httperror.NewHTTPError(http.StatusInternalServerError, "failed to get user by id", err)
	}
	return nil
}

// Full update existing user in DB.
func (r dbRepo) UpdateUser(user *entity.User) error {
	// save user in DB
	updateResult := r.dbClient.Save(user)
	if err := updateResult.Error; err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError, "failed to full update user", err)
	}
	return nil
}
