package mappers

import (
	"github.com/test_server/internal/domain"
	res "github.com/test_server/internal/infra/http/resources"
)

func MapDomainWorkHourToWorkingHoursDto(wh *[]domain.WorkingHours) *[]res.WorkingHoursDTO {
	var result []res.WorkingHoursDTO
	for _, v := range *wh {

		result = append(result, res.WorkingHoursDTO{
			ID:        v.ID,
			WeekDay:   v.WeekDay,
			Date:      v.Date,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return &result
}
func MapWorkHourDtoToWorkHourDomain(wh *[]res.WorkingHoursDTO) *[]domain.WorkingHours {
	var result []domain.WorkingHours
	for _, v := range *wh {
		result = append(result, domain.WorkingHours{
			ID:        v.ID,
			WeekDay:   v.WeekDay,
			Date:      v.Date,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return &result
}
