package app

import (
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type MessageService interface {
	AddMessage(msg *domain.Message) (*domain.Message, error)
	UpdateMessage(msg *domain.Message) (*domain.Message, error)
	PaginateAll(chatID int64, page uint, pageSize uint) (*[]domain.Message, error)
	DeleteMessage(id int64) error
}

type messageService struct {
	repo *database.MessageRepository
}

func NewMessageService(r *database.MessageRepository) MessageService {
	return &messageService{repo: r}
}

func (s *messageService) AddMessage(msg *domain.Message) (*domain.Message, error) {
	return (*s.repo).AddMessage(msg)
}

func (s *messageService) UpdateMessage(msg *domain.Message) (*domain.Message, error) {
	return (*s.repo).UpdateMessage(msg)
}

func (s *messageService) PaginateAll(chatID int64, page uint, pageSize uint) (*[]domain.Message, error) {
	return (*s.repo).PaginateAll(chatID, page, pageSize)
}

func (s *messageService) DeleteMessage(id int64) error {
	return (*s.repo).DeleteMessage(id)
}
