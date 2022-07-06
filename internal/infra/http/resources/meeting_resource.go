package resources

import (
	"time"

	"github.com/test_server/internal/domain"
)

type MeetingDto struct {
	Id             int64     `json:"id"`
	PsychologistId int64     `json:"psychologist_id"`
	ClientId       int64     `json:"client_id"`
	MeetingDate    time.Time `json:"meeting_date"`
	StartTime      float64   `json:"start_time"`
	EndTime        float64   `json:"end_time"`
	Status         string    `json:"status"`
}

func MapDomainToMeetingDto(meeting *domain.Meeting) *MeetingDto {
	return &MeetingDto{
		Id:             meeting.Id,
		PsychologistId: meeting.PsychologistId,
		ClientId:       meeting.ClientId,
		MeetingDate:    meeting.MeetingDate,
		StartTime:      meeting.StartTime,
		EndTime:        meeting.EndTime,
		Status:         meeting.Status,
	}
}

func MapDomainToMeetingDtoCollection(meetings *[]domain.Meeting) *[]MeetingDto {
	var result []MeetingDto
	for _, m := range *meetings {
		dto := MapDomainToMeetingDto(&m)
		result = append(result, *dto)
	}
	return &result
}
