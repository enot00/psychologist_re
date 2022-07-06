package resources

import (
	"time"

	"github.com/test_server/internal/domain"
)

type ClientDto struct {
	Id               int64     `json:"id"`
	UserName         string    `json:"user_name"`
	PhoneNumber      string    `json:"phone_number"`
	Email            string    `json:"email"`
	Avatar           string    `json:"avatar"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registration_date"`
}

func MapDomainToClientDto(client *domain.Client) *ClientDto {
	return &ClientDto{
		Id:               client.Id,
		UserName:         client.UserName,
		PhoneNumber:      client.PhoneNumber,
		Email:            client.Email,
		Avatar:           client.Avatar,
		Password:         client.Password,
		RegistrationDate: client.RegistrationDate,
	}
}

func MapDomainToClientDtoCollection(clients *[]domain.Client) *[]ClientDto {
	var result []ClientDto
	for _, c := range *clients {
		dto := MapDomainToClientDto(&c)
		result = append(result, *dto)
	}
	return &result
}
