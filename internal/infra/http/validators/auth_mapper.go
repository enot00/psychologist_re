package validators

import "github.com/test_server/internal/domain"

type authRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func MapAuthDataToModel(r *authRequest) *domain.User {
	return &domain.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
