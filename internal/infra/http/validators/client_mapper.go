package validators

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/test_server/internal/domain"
)

type clientRequest struct {
	UserName         string    `json:"user_name" validate:"required,gte=2,alphanumunicode"`
	PhoneNumber      string    `json:"phone_number" validate:"required,e164"`
	Email            string    `json:"email" validate:"required,email"`
	Avatar           string    `json:"avatar" validate:"url"`
	OldPassword      string    `json:"old_password"`
	Password         string    `json:"password" validate:"required,custom_validator"`
	RegistrationDate time.Time `json:"registration_date"`
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func mapClientRequestToDomain(clientRequest *clientRequest) *domain.Client {
	var clnt domain.Client
	clnt.UserName = clientRequest.UserName
	clnt.PhoneNumber = clientRequest.PhoneNumber
	clnt.Email = clientRequest.Email
	clnt.Avatar = clientRequest.Avatar
	clnt.OldPassword = clientRequest.OldPassword
	clnt.Password = clientRequest.Password
	clnt.RegistrationDate = time.Now()
	return &clnt
}
