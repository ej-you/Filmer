package usecase

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"
	"Filmer/server/pkg/utils"
	"Filmer/server/config"

	"Filmer/server/internal/auth"
)


// auth.Usecase interface implementation
type authUsecase struct {
	cfg 			*config.Config
	authRepo 		auth.Repository
	authCacheRepo	auth.CacheRepository
}

// auth.Usecase constructor
// Returns auth.Usecase interface
func NewUsecase(cfg *config.Config, authRepo auth.Repository, authCacheRepo auth.CacheRepository) auth.Usecase {
	return &authUsecase{
		cfg: cfg,
		authRepo: authRepo,
		authCacheRepo: authCacheRepo,
	}
}

// Sign up new user
// User email (user.Email) and password (user.Password) must be presented
// Returns *entity.UserWithToken with filled given user struct and random-generated access token
func (this authUsecase) SignUp(user *entity.User) (*entity.UserWithToken, error) {
	// hash password
	passwordHash, err := this.encodePassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}
	user.Password = passwordHash

	// create user
	err = this.authRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	// generate access token
	accessToken, err := utils.ObtainToken(this.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.SignUp: %w", err)
	}

	return &entity.UserWithToken{
		User: user,
		AccessToken: accessToken,
	}, nil
}

// Log in existing user
// User email (user.Email) and password (user.Password) must be presented
// Returns *entity.UserWithToken with filled given user struct and random-generated access token
func (this authUsecase) Login(user *entity.User) (*entity.UserWithToken, error) {
	// password entered by user
	enteredPasswd := user.Password

	// get user from DB with email
	err := this.authRepo.GetUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	// check entered password is correct
	if !this.passwordIsCorrect(enteredPasswd, user.Password) {
		return nil, httpError.NewHTTPError(401, "invalid password", fmt.Errorf("authUsecase.Login"))
	}
	
	// generate access token
	accessToken, err := utils.ObtainToken(this.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("authUsecase.Login: %w", err)
	}

	return &entity.UserWithToken{
		User: user,
		AccessToken: accessToken,
	}, nil
}

// Log out user by set token to blacklist
func (this authUsecase) Logout(token string) error {
	// put token to blacklist
	if err := this.authCacheRepo.SetTokenToBlacklist(token); err != nil {
		return fmt.Errorf("authUsecase.Logout: %w", err)
	}
	return nil
}


// Restrict user access with a blacklisted token
// Return error, if error occurs OR given token in blacklist
func (this authUsecase) RestrictBlacklistedToken(token string) error {
	// search token in blacklist
	isBlacklisted, err := this.authCacheRepo.TokenIsBlacklisted(token)
	if err != nil {
		return fmt.Errorf("authUsecase.RestrictBlacklistedToken: %w", err)
	}
	// return forbidden error if token is in blacklist
	if isBlacklisted {
		return httpError.NewHTTPError(403, "token is not valid", fmt.Errorf("authUsecase.RestrictBlacklistedToken"))
	}
	return nil
}

// Encode given password
// Returns encoded password like a hash
func (this authUsecase) encodePassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		// probably, password is too long
		return nil, httpError.NewHTTPError(400, "failed to encode password", err)
	}
	return hash, nil
}

// Check the given password is equal to its hash from DB
// Returns true, if password is equal
func (this authUsecase) passwordIsCorrect(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
