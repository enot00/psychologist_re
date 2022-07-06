package database

import (
	"errors"
	"log"
	"time"

	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

type client struct {
	Id               int64     `db:"id,omitempty"`
	UserName         string    `db:"user_name"`
	PhoneNumber      string    `db:"phone_number"`
	Email            string    `db:"email"`
	Avatar           string    `db:"avatar"`
	Password         string    `db:"password"`
	RegistrationDate time.Time `db:"registration_date"`
}

type ClientRepository interface {
	SaveClient(client *domain.Client) (*domain.Client, error)
	PaginateAllClients(page uint, pageSize uint) (*[]domain.Client, error)
	FindOneClient(clientId int64) (*domain.Client, error)
	FindByEmail(email string) (*domain.Client, error)
	UpdateClient(client *domain.Client) error
	DeleteClient(clientId int64) error
}

type clientRepository struct {
	coll db.Collection
}

func NewClientRepository(dbSession *db.Session) ClientRepository {
	return &clientRepository{
		coll: (*dbSession).Collection("client"),
	}
}

func (r *clientRepository) SaveClient(client *domain.Client) (*domain.Client, error) {
	clnt := mapDomainToClientDbModel(client)

	err := r.coll.InsertReturning(clnt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapClientDbModelToDomain(clnt), nil
}

func (r *clientRepository) PaginateAllClients(page uint, pageSize uint) (*[]domain.Client, error) {
	var clients []client
	err := r.coll.Find().Paginate(pageSize).Page(page).All(&clients)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapClientDbModelToDomainCollection(&clients), nil
}

func (r *clientRepository) FindOneClient(clientId int64) (*domain.Client, error) {
	var clnt *client
	err := r.coll.Find(db.Cond{"id": clientId}).One(&clnt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapClientDbModelToDomain(clnt), nil
}

func (r *clientRepository) FindByEmail(email string) (*domain.Client, error) {
	var c client

	err := r.coll.Find(db.Cond{"email": email}).One(&c)
	if err != nil {
		log.Printf("db error ClientRepository.FindOne(): %v", err)
		return nil, err
	}

	return mapClientDbModelToDomain(&c), nil
}

func (r *clientRepository) UpdateClient(domClient *domain.Client) error {
	var dbClient client
	res := r.coll.Find(db.Cond{"id": domClient.Id})
	err := res.One(&dbClient)
	if err != nil {
		log.Print(err)
		return err
	}

	// to update password, you must confirm the old password
	err = bcrypt.CompareHashAndPassword([]byte(dbClient.Password), []byte(domClient.OldPassword))
	if err != nil {
		err = errors.New("old password is wrong")
		log.Print(err)
		return err
	}

	updClnt := mapDomainToClientDbModel(domClient)
	// RegistrationDate should not be updated
	updClnt.RegistrationDate = dbClient.RegistrationDate
	err = res.Update(updClnt)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (r *clientRepository) DeleteClient(clientId int64) error {
	var clnt *client
	res := r.coll.Find(clientId)

	err := res.One(&clnt)
	if err != nil {
		log.Print(err)
		return err
	}

	err = res.Delete()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func mapClientDbModelToDomain(client *client) *domain.Client {
	return &domain.Client{
		Id:               client.Id,
		UserName:         client.UserName,
		PhoneNumber:      client.PhoneNumber,
		Email:            client.Email,
		Avatar:           client.Avatar,
		Password:         client.Password,
		RegistrationDate: client.RegistrationDate,
	}
}

func mapClientDbModelToDomainCollection(clients *[]client) *[]domain.Client {
	var result []domain.Client

	for _, cl := range *clients {
		newCl := mapClientDbModelToDomain(&cl)
		result = append(result, *newCl)
	}
	return &result
}

func mapDomainToClientDbModel(clnt *domain.Client) *client {
	return &client{
		Id:               clnt.Id,
		UserName:         clnt.UserName,
		PhoneNumber:      clnt.PhoneNumber,
		Email:            clnt.Email,
		Avatar:           clnt.Avatar,
		Password:         clnt.Password,
		RegistrationDate: clnt.RegistrationDate,
	}
}
