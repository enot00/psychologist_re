package app

import (
	"log"

	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type ClientService interface {
	SaveClient(client *domain.Client) (*domain.Client, error)
	PaginateAllClients(page uint, pageSize uint) (*[]domain.Client, error)
	FindOneClient(clientId int64) (*domain.Client, error)
	UpdateClient(client *domain.Client) error
	DeleteClient(clientId int64) error
}

type clientService struct {
	clientRepo *database.ClientRepository
}

func NewClientService(c *database.ClientRepository) ClientService {
	return &clientService{
		clientRepo: c,
	}
}

func (s *clientService) SaveClient(client *domain.Client) (*domain.Client, error) {
	//client.Password = helpers.HashPassword(client.Password)
	//
	//client, err := (*s.clientRepo).SaveClient(client)
	//if err != nil {
	//	log.Print(err)
	//	return nil, err
	//}

	return nil, nil
}

func (s *clientService) PaginateAllClients(page uint, pageSize uint) (*[]domain.Client, error) {
	clients, err := (*s.clientRepo).PaginateAllClients(page, pageSize)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return clients, nil
}

func (s *clientService) FindOneClient(clientId int64) (*domain.Client, error) {
	client, err := (*s.clientRepo).FindOneClient(clientId)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return client, nil
}

func (s *clientService) UpdateClient(client *domain.Client) error {
	err := (*s.clientRepo).UpdateClient(client)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s *clientService) DeleteClient(clientId int64) error {
	err := (*s.clientRepo).DeleteClient(clientId)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
