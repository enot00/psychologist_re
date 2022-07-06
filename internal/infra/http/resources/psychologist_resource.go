package resources

import "github.com/test_server/internal/domain"

type PsychologistDTO struct {
	User           UserDTO             `json:"user"`
	Description    string              `json:"description"`
	Specialization []SpecializationDTO `json:"specialization"`
}
type PsychologistDTOShort struct {
	ID             int64               `json:"id"`
	User           UserDTOShort        `json:"user"`
	Specialization []SpecializationDTO `json:"specialization"`
}

type SpecializationDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func MapPsychologistDomainToDto(psyDtoShort *domain.Psychologist) PsychologistDTO {
	spec := mapDomainSpecializationToDtoSpecialization(&psyDtoShort.Specialization)

	return PsychologistDTO{
		User:           *MapUserDomainToDto(&psyDtoShort.User),
		Description:    psyDtoShort.Description,
		Specialization: *spec,
	}
}
func mapDomainToDtoShort(psyDtoShort *domain.Psychologist) PsychologistDTOShort {
	spec := mapDomainSpecializationToDtoSpecialization(&psyDtoShort.Specialization)

	return PsychologistDTOShort{
		ID: psyDtoShort.User.ID,
		User: UserDTOShort{
			FirstName:  psyDtoShort.User.FirstName,
			SecondName: psyDtoShort.User.SecondName,
			LastName:   psyDtoShort.User.LastName,
			Avatar:     psyDtoShort.User.GetAvatar(),
		},
		Specialization: *spec,
	}
}
func mapDomainSpecializationToDtoSpecialization(sp *[]domain.Specialization) *[]SpecializationDTO {
	var result []SpecializationDTO
	for _, v := range *sp {
		result = append(result, SpecializationDTO{ID: v.ID, Name: v.Name})
	}
	return &result
}

func MapPsychologistDomainToDtoCollection(psy *[]domain.Psychologist) *[]PsychologistDTOShort {
	var result []PsychologistDTOShort
	for _, c := range *psy {
		dto := mapDomainToDtoShort(&c)
		result = append(result, dto)
	}
	return &result
}
