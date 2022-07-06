package app

import (
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type UserService interface {
	PaginateAll(page, pageSize uint) (*[]domain.UserShort, error)
	GetOneByID(id int64) (*domain.User, error)
	GetOneByEmail(email string) (*domain.User, error)
	Save(user *domain.User) (*domain.User, error)
	Update(user *domain.User) error
	ResetPassword(id int64, newPassword string) error
	Delete(userID int64) error
}

type userService struct {
	repo *database.UserRepository
	hs   *HashService
}

func NewUserService(repo *database.UserRepository, hs *HashService) UserService {
	return &userService{repo: repo, hs: hs}
}

func (s *userService) PaginateAll(page, pageSize uint) (*[]domain.UserShort, error) {
	return (*s.repo).PaginateAll(page, pageSize)
}
func (s *userService) GetOneByID(id int64) (*domain.User, error) {
	return (*s.repo).GetOneByID(id)
}
func (s *userService) GetOneByEmail(email string) (*domain.User, error) {
	return (*s.repo).GetOneByEmail(email)
}
func (s *userService) Save(user *domain.User) (*domain.User, error) {
	user.Password = (*s.hs).Hash(user.Password)
	return (*s.repo).Save(user)
}

func (s *userService) Update(user *domain.User) error {
	return (*s.repo).Update(user)
}

func (s *userService) ResetPassword(id int64, newPassword string) error {
	err := (*s.repo).Update(&domain.User{
		ID:       id,
		Password: (*s.hs).Hash(newPassword),
	})

	return err
}
func (s *userService) Delete(id int64) error {
	return (*s.repo).Delete(id)
}
