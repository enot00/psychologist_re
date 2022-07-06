package app

import (
	"github.com/test_server/internal/infra/database"
	res "github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/resources/mappers"
)

type WorkingHoursService interface {
	GetOneByID(id int64) (*[]res.WorkingHoursDTO, error)
	Save(workingHours *[]res.WorkingHoursDTO) (*[]res.WorkingHoursDTO, error)
	Update(id int64, dto *[]res.WorkingHoursDTO) error
	Delete(id int64) error
}
type workingHoursService struct {
	repo *database.WorkingHoursRepository
}

func NewWorkingHoursService(d *database.WorkingHoursRepository) WorkingHoursService {
	return &workingHoursService{
		repo: d,
	}
}
func (w *workingHoursService) GetOneByID(id int64) (*[]res.WorkingHoursDTO, error) {
	getOneByID, err := (*w.repo).GetOneByID(id)
	if err != nil {
		return nil, err
	}
	return mappers.MapDomainWorkHourToWorkingHoursDto(getOneByID), nil
}
func (w *workingHoursService) Save(wh *[]res.WorkingHoursDTO) (*[]res.WorkingHoursDTO, error) {
	whSave, err := (*w.repo).Save(mappers.MapWorkHourDtoToWorkHourDomain(wh))
	if err != nil {
		return nil, err
	}
	return mappers.MapDomainWorkHourToWorkingHoursDto(whSave), nil
}
func (w *workingHoursService) Update(id int64, wh *[]res.WorkingHoursDTO) error {
	err := (*w.repo).Update(id, mappers.MapWorkHourDtoToWorkHourDomain(wh))
	if err != nil {
		return err
	}
	return nil
}
func (w *workingHoursService) Delete(id int64) error {
	err := (*w.repo).Delete(id)
	if err != nil {
		return err
	}
	return nil
}
