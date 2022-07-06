package app

import (
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type PsychologistService interface {
	PaginateAll(page, pageSize uint) (*[]domain.Psychologist, error)
	GetOne(id uint) (*domain.Psychologist, error)
	Update(psy *domain.Psychologist) error
}

type psychologistService struct {
	repo *database.PsychologistRepository
}

func NewPsychologistService(d *database.PsychologistRepository) PsychologistService {
	return &psychologistService{
		repo: d,
	}
}

func (s *psychologistService) PaginateAll(page, pageSize uint) (*[]domain.Psychologist, error) {
	return (*s.repo).PaginateAll(page, pageSize)
}
func (s *psychologistService) GetOne(id uint) (*domain.Psychologist, error) {
	return (*s.repo).GetOne(id)
}

func (s *psychologistService) Update(psychologists *domain.Psychologist) error {
	return (*s.repo).Update(psychologists)
}
