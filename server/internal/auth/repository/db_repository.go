package repository

import (
	"errors"
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
func (this authRepository) CreateUser(user *entity.User) error {
	createResult := this.dbClient.Create(user)
	if err := createResult.Error; err != nil {
		// if user with such email already exists
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return httpError.NewHTTPError(409, "user with such email already exists", err)
		}
		return httpError.NewHTTPError(500, "failed to create user", err)
	}
	return nil
}

// Get user by email
// Fill given user struct
// Returns error even if user not found
func (this authRepository) GetUserByEmail(user *entity.User) error {
	selectResult := this.dbClient.Where("email = ?", user.Email).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httpError.NewHTTPError(404, "user with such email was not found", err)
		}
		return httpError.NewHTTPError(500, "failed to get user by email", err)
	}
	return nil
}
