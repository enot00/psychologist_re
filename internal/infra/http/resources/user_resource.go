package resources

import (
	"github.com/test_server/internal/domain"
	"time"
)

type UserDTO struct {
	ID         int64           `json:"id"`
	Phone      string          `json:"phone"`
	Email      string          `json:"email,omitempty"`
	Password   string          `json:"password"`
	FirstName  string          `json:"first_name"`
	SecondName *string         `json:"second_name,omitempty"`
	LastName   string          `json:"last_name"`
	Avatar     string          `json:"avatar"`
	Role       domain.UserRole `json:"role"`
	CreatedAt  time.Time       `json:"created_at"`
	DeletedAt  time.Time       `json:"deleted_at,omitempty"`
}

type UserDTOShort struct {
	FirstName  string  `json:"first_name"`
	SecondName *string `json:"second_name,omitempty"`
	LastName   string  `json:"last_name"`
	Avatar     string  `json:"avatar"`
}

type UserResource struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func MapModelToUserResource(c *domain.User) *UserResource {
	return &UserResource{
		Id:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}
}
func MapUserDomainToDto(user *domain.User) *UserDTO {

	return &UserDTO{
		ID:         user.ID,
		Phone:      user.Phone,
		Email:      user.Email,
		Password:   user.Password,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		LastName:   user.LastName,
		Avatar:     user.GetAvatar(),
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		DeletedAt:  user.DeletedAt,
	}
}
func MapUserShortDomainToDto(userShort *domain.UserShort) *UserDTOShort {

	return &UserDTOShort{
		FirstName:  userShort.FirstName,
		SecondName: userShort.SecondName,
		LastName:   userShort.LastName,
		Avatar:     userShort.GetAvatarShort(),
	}
}

func MapUserDomainToDtoCollection(users *[]domain.UserShort) *[]UserDTOShort {
	var result []UserDTOShort
	for _, v := range *users {
		dto := MapUserShortDomainToDto(&v)
		result = append(result, *dto)
	}
	return &result
}
