package repository

import (
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"

	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"

	"Filmer/server/internal/auth"
)

// auth.Repository interface implementation
type authRepository struct {
	dbClient *gorm.DB
}

// auth.Repository constructor
// Returns auth.Repository interface
func NewRepository(dbClient *gorm.DB) auth.Repository {
	return &authRepository{
		dbClient: dbClient,
	}
}

// Create new user
// Fill given user struct
func (ar authRepository) CreateUser(user *entity.User) error {
	createResult := ar.dbClient.Create(user)
	if err := createResult.Error; err != nil {
		// if user with such email already exists
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return httpError.NewHTTPError(http.StatusConflict, "user with such email already exists", err)
		}
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to create user", err)
	}
	return nil
}

// Get user by email
// User email (user.Email) must be presented
// Fill given user struct
// Returns error even if user not found
func (ar authRepository) GetUserByEmail(user *entity.User) error {
	selectResult := ar.dbClient.Where("email = ?", user.Email).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httpError.NewHTTPError(http.StatusNotFound, "user with such email was not found", err)
		}
		return httpError.NewHTTPError(http.StatusInternalServerError, "failed to get user by email", err)
	}
	return nil
}
