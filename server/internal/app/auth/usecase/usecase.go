package usecase

import (
	"fmt"
	"net/http"

	"Filmer/server/config"
	"Filmer/server/internal/app/entity"
	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/password"
	"Filmer/server/internal/pkg/token"

	"Filmer/server/internal/app/auth"
)

var _ auth.Usecase = (*usecase)(nil)

// auth.Usecase implementation.
type usecase struct {
	cfg           *config.Config
	authDBRepo    auth.DBRepo
	authCacheRepo auth.CacheRepo
}

// Returns auth.Usecase interface.
func NewUsecase(cfg *config.Config,
	authDBRepo auth.DBRepo, authCacheRepo auth.CacheRepo) auth.Usecase {

	return &usecase{
		cfg:           cfg,
		authDBRepo:    authDBRepo,
		authCacheRepo: authCacheRepo,
	}
}

// Sign up new user.
// User email (user.Email) and password (user.Password) must be presented.
// Returns *entity.UserWithToken with filled given user struct and random-generated access token.
func (u usecase) SignUp(user *entity.User) (*entity.UserWithToken, error) {
	// hash password
	passwordHash, err := password.Encode(user.Password)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}
	user.Password = passwordHash

	// create user
	err = u.authDBRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	// generate access token
	accessToken, err := token.New(u.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	return &entity.UserWithToken{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

// Log in existing user.
// User email (user.Email) and password (user.Password) must be presented.
// Returns *entity.UserWithToken with filled given user struct and random-generated access token.
func (u usecase) Login(user *entity.User) (*entity.UserWithToken, error) {
	// password entered by user
	enteredPasswd := user.Password

	// get user from DB with email
	err := u.authDBRepo.GetUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	// check entered password is correct
	if !password.IsCorrect(enteredPasswd, user.Password) {
		return nil, httperror.New(http.StatusUnauthorized,
			"invalid password", fmt.Errorf("authUsecase.Login"))
	}

	// generate access token
	accessToken, err := token.New(u.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	return &entity.UserWithToken{
		User:        user,
		AccessToken: accessToken,
	}, nil
}

// Log out user by set token to blacklist.
func (u usecase) Logout(token string) error {
	// put token to blacklist
	if err := u.authCacheRepo.SetTokenToBlacklist(token); err != nil {
		return fmt.Errorf("authUsecase.Logout: %w", err)
	}
	return nil
}

// Restrict user access with a blacklisted token.
// Return error, if error occurs OR given token in blacklist.
func (u usecase) RestrictBlacklistedToken(token string) error {
	// search token in blacklist
	isBlacklisted, err := u.authCacheRepo.TokenIsBlacklisted(token)
	if err != nil {
		return fmt.Errorf("authUsecase.RestrictBlacklistedToken: %w", err)
	}
	// return forbidden error if token is in blacklist
	if isBlacklisted {
		return httperror.New(http.StatusForbidden,
			"token is not valid", fmt.Errorf("authUsecase.RestrictBlacklistedToken"))
	}
	return nil
}
