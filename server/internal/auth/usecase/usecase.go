package usecase

import (
	"fmt"
	"net/http"

	"Filmer/server/config"
	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/utils"

	"Filmer/server/internal/auth"
)

// auth.Usecase interface implementation
type authUsecase struct {
	cfg           *config.Config
	authRepo      auth.Repository
	authCacheRepo auth.CacheRepository
}

// auth.Usecase constructor
// Returns auth.Usecase interface
func NewUsecase(cfg *config.Config, authRepo auth.Repository, authCacheRepo auth.CacheRepository) auth.Usecase {
	return &authUsecase{
		cfg:           cfg,
		authRepo:      authRepo,
		authCacheRepo: authCacheRepo,
	}
}

// Sign up new user
// User email (user.Email) and password (user.Password) must be presented
// Returns *entity.UserWithToken with filled given user struct and random-generated access token
func (au authUsecase) SignUp(user *entity.User) (*entity.UserWithToken, error) {
	// hash password
	passwordHash, err := utils.EncodePassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}
	user.Password = passwordHash

	// create user
	err = au.authRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	// generate access token
	accessToken, err := utils.ObtainToken(au.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	return &entity.UserWithToken{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

// Log in existing user
// User email (user.Email) and password (user.Password) must be presented
// Returns *entity.UserWithToken with filled given user struct and random-generated access token
func (au authUsecase) Login(user *entity.User) (*entity.UserWithToken, error) {
	// password entered by user
	enteredPasswd := user.Password

	// get user from DB with email
	err := au.authRepo.GetUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	// check entered password is correct
	if !utils.PasswordIsCorrect(enteredPasswd, user.Password) {
		return nil, httpError.NewHTTPError(http.StatusUnauthorized, "invalid password", fmt.Errorf("authUsecase.Login"))
	}

	// generate access token
	accessToken, err := utils.ObtainToken(au.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	return &entity.UserWithToken{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

// Log out user by set token to blacklist
func (au authUsecase) Logout(token string) error {
	// put token to blacklist
	if err := au.authCacheRepo.SetTokenToBlacklist(token); err != nil {
		return fmt.Errorf("authUsecase.Logout: %w", err)
	}
	return nil
}

// Restrict user access with a blacklisted token
// Return error, if error occurs OR given token in blacklist
func (au authUsecase) RestrictBlacklistedToken(token string) error {
	// search token in blacklist
	isBlacklisted, err := au.authCacheRepo.TokenIsBlacklisted(token)
	if err != nil {
		return fmt.Errorf("authUsecase.RestrictBlacklistedToken: %w", err)
	}
	// return forbidden error if token is in blacklist
	if isBlacklisted {
		return httpError.NewHTTPError(http.StatusForbidden, "token is not valid", fmt.Errorf("authUsecase.RestrictBlacklistedToken"))
	}
	return nil
}
