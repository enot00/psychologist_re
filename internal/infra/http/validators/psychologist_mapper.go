package validators

import (
	"github.com/test_server/internal/domain"
)

type PsychologistRequest struct {
	User           userUpdateRequest       `json:"user"`
	Description    string                  `json:"description"`
	Specialization []SpecializationRequest `json:"specialization"`
}

type SpecializationRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type psychologistAllPageRequest struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page-size" validate:"required"`
}
type psychologistIDRequest struct {
	ID int64 `json:"id" validate:"required"`
}

func MapPsychologistRequestToDomain(psyRequest *PsychologistRequest) *domain.Psychologist {
	return &domain.Psychologist{
		User:           *MapUserRequestUpdateToDomain(&psyRequest.User),
		Description:    psyRequest.Description,
		Specialization: *mapSpecializationRequestToDomain(&psyRequest.Specialization),
	}
}
func mapSpecializationRequestToDomain(sp *[]SpecializationRequest) *[]domain.Specialization {
	var result []domain.Specialization
	for _, v := range *sp {
		result = append(result, domain.Specialization{ID: v.ID, Name: v.Name})
	}
	return &result
}
func MapPsychologistRequestAllPage(page, pageSize int) *psychologistAllPageRequest {
	return &psychologistAllPageRequest{
		Page:     page,
		PageSize: pageSize,
	}
}
func MapPsychologistIDRequest(id int64) *psychologistIDRequest {
	return &psychologistIDRequest{ID: id}
}
