package validators

import (
	"github.com/test_server/internal/domain"
)

type userUpdateRequest struct {
	ID         int64   `json:"id" validate:"required"`
	Phone      string  `json:"phone,omitempty"`
	Email      string  `json:"email,omitempty"`
	FirstName  string  `json:"first_name,omitempty"`
	SecondName *string `json:"second_name,omitempty"`
	LastName   string  `json:"last_name,omitempty"`
	Avatar     string  `json:"avatar,omitempty"`
}

type newUserRequest struct {
	FirstName  string  `json:"first_name" validate:"required"`
	SecondName *string `json:"second_name,omitempty"`
	LastName   string  `json:"last_name" validate:"required"`
	Phone      string  `json:"phone" validate:"required"`
	Email      string  `json:"email" validate:"required"`
	Password   string  `json:"password" validate:"required"`
}

type userAllPageRequest struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page-size" validate:"required"`
}

type userIDRequest struct {
	ID int64 `json:"id" validate:"required"`
}

func MapUserRequestUpdateToDomain(userRequest *userUpdateRequest) *domain.User {
	return &domain.User{
		ID:         userRequest.ID,
		Phone:      userRequest.Phone,
		Email:      userRequest.Email,
		FirstName:  userRequest.FirstName,
		SecondName: userRequest.SecondName,
		LastName:   userRequest.LastName,
		Avatar:     userRequest.Avatar,
	}
}

func MapNewUserToModel(r *newUserRequest) *domain.User {
	return &domain.User{
		FirstName:  r.FirstName,
		SecondName: r.SecondName,
		LastName:   r.LastName,
		Phone:      r.Phone,
		Email:      r.Email,
		Password:   r.Password,
	}
}
func MapUserRequestAllPage(page, pageSize int) *userAllPageRequest {
	return &userAllPageRequest{
		Page:     page,
		PageSize: pageSize,
	}
}
func MapUserIDRequest(id int64) *userIDRequest {
	return &userIDRequest{ID: id}
}
