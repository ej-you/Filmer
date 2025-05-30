package usecase

import (
	"bytes"
	"fmt"
	"net/http"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/app/user"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/password"
)

var _ user.Usecase = (*usecase)(nil)

// user.Usecase implementation.
type usecase struct {
	cfg        *config.Config
	userDBRepo user.DBRepo
}

// Returns user.Usecase interface.
func NewUsecase(cfg *config.Config, userDBRepo user.DBRepo) user.Usecase {
	return &usecase{
		cfg:        cfg,
		userDBRepo: userDBRepo,
	}
}

// Change user password.
// User ID (user.ID) and entered current password (user.Password) must be presented.
func (u usecase) ChangePassword(user *entity.User, newPassword []byte) error {
	// password entered by user
	enteredPasswd := user.Password

	// get user by ID
	err := u.userDBRepo.GetUserByID(user)
	if err != nil {
		return fmt.Errorf("user usecase.ChangePassword: %w", err)
	}
	// check entered password is correct
	if !password.IsCorrect(enteredPasswd, user.Password) {
		return httperror.New(http.StatusBadRequest,
			"invalid current password", fmt.Errorf("user usecase.ChangePassword"))
	}
	// if a newPassword is equal to the current user password
	if bytes.Equal(newPassword, enteredPasswd) {
		return httperror.New(http.StatusBadRequest,
			"cannot use the current password as a new password",
			fmt.Errorf("user usecase.ChangePassword"))
	}

	// hash new user password
	user.Password, err = password.Encode(newPassword)
	if err != nil {
		return fmt.Errorf("user usecase.ChangePassword: %w", err)
	}

	// update user
	err = u.userDBRepo.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("user usecase.ChangePassword: %w", err)
	}
	return nil
}
