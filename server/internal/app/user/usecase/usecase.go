package usecase

import (
	"bytes"
	"fmt"
	"net/http"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	httpError "Filmer/server/internal/pkg/http_error"
	"Filmer/server/internal/pkg/utils"

	"Filmer/server/internal/app/user"
)

// user.Usecase interface implementation
type userUsecase struct {
	cfg      *config.Config
	userRepo user.Repository
}

// user.Usecase constructor
// Returns user.Usecase interface
func NewUsecase(cfg *config.Config, userRepo user.Repository) user.Usecase {
	return &userUsecase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

// Change user password
// User ID (user.ID) and entered current password (user.Password) must be presented
func (uu userUsecase) ChangePassword(user *entity.User, newPassword []byte) error {
	// password entered by user
	enteredPasswd := user.Password

	// get user by ID
	err := uu.userRepo.GetUserByID(user)
	if err != nil {
		return fmt.Errorf("userUsecase.ChangePassword: %w", err)
	}
	// check entered password is correct
	if !utils.PasswordIsCorrect(enteredPasswd, user.Password) {
		return httpError.NewHTTPError(http.StatusBadRequest, "invalid current password", fmt.Errorf("userUsecase.ChangePassword"))
	}
	// if a newPassword is equal to the current user password
	if bytes.Equal(newPassword, enteredPasswd) {
		return httpError.NewHTTPError(http.StatusBadRequest, "cannot use the current password as a new password", fmt.Errorf("userUsecase.ChangePassword"))
	}

	// hash new user password
	user.Password, err = utils.EncodePassword(newPassword)
	if err != nil {
		return fmt.Errorf("userUsecase.ChangePassword: %w", err)
	}

	// update user
	err = uu.userRepo.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("userUsecase.ChangePassword: %w", err)
	}

	return nil
}
