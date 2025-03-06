package repository

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"
	
	"Filmer/server/internal/auth"
)


// тип интерфейса auth.Repository
type authRepository struct {
	dbClient *gorm.DB
}

// конструктор для типа интерфейса auth.Repository
func NewRepository(dbClient *gorm.DB) auth.Repository {
	return &authRepository{
		dbClient: dbClient,
	}
}

// создание нового юзера в БД
func (this authRepository) CreateUser(user *entity.User) (*entity.User, error) {
	createResult := this.dbClient.Create(user)
	if err := createResult.Error; err != nil {
		// если юзер с таким юзернеймом уже есть
		if strings.HasSuffix(err.Error(), "(SQLSTATE 23505)") {
			return nil, fmt.Errorf("create user: %w", httpError.NewHTTPError(409, "user with such email already exists"))
		}
		return nil, httpError.NewHTTPError(500, "failed to create user: " + err.Error())
	}
	return user, nil
}

// получение юзера из БД по email
func (this authRepository) FindUserByEmail(user *entity.User) (*entity.User, error) {
	selectResult := this.dbClient.Where("email = ?", user.Email).First(user)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("fing user by email: %w", httpError.NewHTTPError(404, "user with such email was not found"))
		}
		return nil, httpError.NewHTTPError(500, "failed to fing user by email: " + err.Error())
	}
	return user, nil
}
