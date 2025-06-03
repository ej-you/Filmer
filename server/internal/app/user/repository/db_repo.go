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

// GetUserByID gets user by ID.
// User ID (user.ID) must be presented.
// Fill given user struct.
// Returns error even if user not found.
func (r dbRepo) GetUserByID(user *entity.User) error {
	selectResult := r.dbClient.Where("id = ?", user.ID).First(user)
	if err := selectResult.Error; err != nil {
		// if user nof found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return httperror.New(http.StatusNotFound, "user with such id was not found", err)
		}
		return httperror.New(http.StatusInternalServerError, "failed to get user by id", err)
	}
	return nil
}

// UpdateUser fully updates existing user in DB.
func (r dbRepo) UpdateUser(user *entity.User) error {
	// save user in DB
	updateResult := r.dbClient.Save(user)
	if err := updateResult.Error; err != nil {
		return httperror.New(http.StatusInternalServerError, "failed to full update user", err)
	}
	return nil
}

// GetActivity gets amount of user movies (for every user) grouped by his lists
// such as "stared", "want" and "watched".
func (r dbRepo) GetActivity() (entity.UsersActivity, error) {
	var activity entity.UsersActivity

	err := r.dbClient.Table("users").
		Select(`users.email, users.created_at,
			COUNT(CASE WHEN um.status = 1 THEN 1 END) AS want,
			COUNT(CASE WHEN um.status = 2 THEN 1 END) AS watched,
			COUNT(CASE WHEN um.stared = true THEN 1 END) AS stared`).
		Joins("LEFT JOIN user_movies AS um ON users.id=um.user_id").
		Group("users.id").
		Scan(&activity).Error
	if err != nil {
		return nil, httperror.New(http.StatusInternalServerError, "get users activity", err)
	}
	return activity, nil
}
