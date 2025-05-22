package repository

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"Filmer/server/internal/app/entity"
	httpError "Filmer/server/internal/pkg/http_error"

	"Filmer/server/internal/app/user"
)

// user.Repository interface implementation
type userRepository struct {
	dbClient *gorm.DB
}

// user.Repository constructor
// Returns user.Repository interface
func NewRepository(dbClient *gorm.DB) user.Repository {
	return &userRepository{
		dbClient: dbClient,
	}
}

// Get user by ID
// User ID (user.ID) must be presented
// Fill given user struct
// Returns error even if user not found
func (ur userRepository) GetUserByID(user *entity.User) error {
	selectResult := ur.dbClient.Where("id = ?", user.ID).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httpError.NewHTTPError(http.StatusNotFound, "user with such id was not found", err)
		}
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to get user by id", err)
	}
	return nil
}

// Full update existing user in DB
func (ur userRepository) UpdateUser(user *entity.User) error {
	// save user in DB
	updateResult := ur.dbClient.Save(user)
	if err := updateResult.Error; err != nil {
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to full update user", err)
	}
	return nil
}
