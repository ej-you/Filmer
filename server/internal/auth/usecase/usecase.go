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


// тип интерфейса auth.Usecase
type authUsecase struct {
	cfg 			*config.Config
	authRepo 		auth.Repository
	authCacheRepo	auth.CacheRepository
}

// конструктор для типа интерфейса auth.Usecase
func NewUsecase(cfg *config.Config, authRepo auth.Repository, authCacheRepo auth.CacheRepository) auth.Usecase {
	return &authUsecase{
		cfg: cfg,
		authRepo: authRepo,
		authCacheRepo: authCacheRepo,
	}
}

func (this authUsecase) SignUp(user *entity.User) (*entity.UserWithToken, error) {
	// хэширование пароля и заполнение структуры
	passwordHash, err := this.encodePassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("sign up: %w", err)
	}
	user.Password = passwordHash

	// создание юзера
	createdUser, err := this.authRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("sign up: %w", err)
	}

	// генерация access токена
	accessToken, err := utils.ObtainToken(this.cfg, createdUser.ID)
	if err != nil {
		return nil, fmt.Errorf("sign up: %w", err)
	}

	return &entity.UserWithToken{
		User: createdUser,
		AccessToken: accessToken,
	}, nil
}

func (this authUsecase) Login(user *entity.User) (*entity.UserWithToken, error) {
	// введённый юзером пароль
	enteredPasswd := user.Password

	// получение юзера из БД по email
	existsUser, err := this.authRepo.FindUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	// сверка введённого пароля с хэшем из БД
	if !this.passwordIsCorrect(enteredPasswd, existsUser.Password) {
		return nil, fmt.Errorf("login: %w", httpError.NewHTTPError(401, "invalid password"))
	}
	
	// генерация access токена
	accessToken, err := utils.ObtainToken(this.cfg, user.ID)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	return &entity.UserWithToken{
		User: existsUser,
		AccessToken: accessToken,
	}, nil
}

func (this authUsecase) Logout(token string) error {
	// помещение токена в чёрный список
	if err := this.authCacheRepo.SetTokenToBlacklist(token); err != nil {
		return fmt.Errorf("logout: %w", err)
	}
	return nil
}

func (this authUsecase) FindUserByEmail(user *entity.User) (*entity.User, error) {
	// получение юзера из БД по email
	foundUser, err := this.authRepo.FindUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	return foundUser, nil
}

func (this authUsecase) RestrictBlacklistedToken(token string) error {
	// поиск токена в черном списке
	isBlacklisted, err := this.authCacheRepo.TokenIsBlacklisted(token)
	if err != nil {
		return fmt.Errorf("check blacklisted token: %w", err)
	}
	// если токен в чёрном списке, то возвращаем ошибку на запрет доступа
	if isBlacklisted {
		return httpError.NewHTTPError(403, "token is not valid")
	}
	return nil
}


// кодирование пароля в хэш
func (this authUsecase) encodePassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		// скорее всего, ошибка из-за слишком длинного пароля
		return nil, fmt.Errorf("failed to encode password: %w", err)
	}
	return hash, nil
}

// проверка введённого юзером пароля на совпадение с хэшем из БД
func (this authUsecase) passwordIsCorrect(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
