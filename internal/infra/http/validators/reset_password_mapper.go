package validators

import "github.com/test_server/internal/domain"

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password" validate:"required"`
}

func MapResetPasswordRequestToModel(r *ResetPasswordRequest) *domain.User {
	return &domain.User{
		Password: r.NewPassword,
	}
}
