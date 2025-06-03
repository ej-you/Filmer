package usecase

import (
	"fmt"

	"Filmer/admin/internal/app/entity"
	"Filmer/admin/internal/app/repo"
)

var _ UserUsecase = (*userUsecase)(nil)

// UserUsecase implementation.
type userUsecase struct {
	apiRepo repo.UserAPIRepo
}

func NewUserUsecase(apiRepo repo.UserAPIRepo) UserUsecase {
	return &userUsecase{
		apiRepo: apiRepo,
	}
}

// GetUsersActivity gets slice of users activity and prepares it to representation.
func (u *userUsecase) GetUsersActivity() (entity.UsersActivity, error) {
	usersData, err := u.apiRepo.GetUsersActivity()
	if err != nil {
		return nil, fmt.Errorf("get users activity: %w", err)
	}
	// format datetime to pretty view
	for userIdx := range usersData {
		usersData[userIdx].FormatCreatedAt()
	}
	return usersData, nil
}
